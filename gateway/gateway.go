package gateway

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Frontman-Labs/frontman/config"
	"github.com/Frontman-Labs/frontman/log"
	"github.com/Frontman-Labs/frontman/plugins"
	"github.com/Frontman-Labs/frontman/service"
)

type APIGateway struct {
	reg   service.ServiceRegistry
	plugs []plugins.FrontmanPlugin
	conf  *config.Config
	log   log.Logger
}

func NewAPIGateway(bs service.ServiceRegistry, plugs []plugins.FrontmanPlugin, conf *config.Config, logger log.Logger) *APIGateway {
	return &APIGateway{
		reg:   bs,
		plugs: plugs,
		conf:  conf,
		log:   logger,
	}
}

func (g *APIGateway) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, plugin := range g.plugs {
		if err := plugin.PreRequest(req, g.reg, g.conf); err != nil {
			g.log.Errorf("Plugin error: %v", err)
			// http.Error(w, err.Error(), err.StatusCode())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(err.StatusCode())
			json.NewEncoder(w).Encode(err.Response())
			return
		}
	}

	// Find the backend service that matches the request
	backendService := g.reg.GetTrie().FindBackendService(req)

	// If the backend service was not found, return a 404 error
	if backendService == nil {
		http.NotFound(w, req)
		return
	}

	// Get the upstream target URL for this request
	upstreamTarget := backendService.GetLoadBalancer().ChooseTarget(backendService.UpstreamTargets)

	urlPath := req.URL.Path
	if backendService.StripPath {
		urlPath = strings.TrimPrefix(req.URL.Path, backendService.Path)
	}

	// Use the compiledRewriteMatch field in the backendService struct to apply the rewrite
	if backendService.GetCompiledRewriteMatch() != nil {
		urlPath = backendService.GetCompiledRewriteMatch().ReplaceAllString(urlPath, backendService.RewriteReplace)
	}

	// Create a new target URL with the service path and scheme
	targetURL, err := url.Parse(upstreamTarget + urlPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Add query parameters if they are available
	if req.URL.RawQuery != "" {
		targetURL.RawQuery = req.URL.RawQuery
	}

	// Get client for backend service
	client := backendService.GetHttpClient()

	// Copy the headers from the original request
	headers := make(http.Header)
	copyHeaders(headers, req.Header)

	if backendService.AuthConfig != nil {
		tokenValidator := backendService.GetTokenValidator()
		// Backend service has auth config specified
		claims, err := tokenValidator.ValidateToken(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims != nil {
			data, err := json.Marshal(claims)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			headers.Add(backendService.GetUserDataHeader(), string(data))
		}

	}
	// Remove the X-Forwarded-For header to prevent spoofing
	headers.Del("X-Forwarded-For")

	// Log a message indicating that the request is being sent to the target service
	g.log.Infof("Sending request to %s: %s %s", upstreamTarget, req.Method, urlPath)

	// Send the request to the target service using the client with the specified transport
	resp, err := client.Do(&http.Request{
		Method:        req.Method,
		URL:           targetURL,
		Proto:         req.Proto,
		ProtoMajor:    req.ProtoMajor,
		ProtoMinor:    req.ProtoMinor,
		Header:        headers,
		Body:          req.Body,
		ContentLength: req.ContentLength,
		Host:          targetURL.Host,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		g.log.Infof("Error sending request: %v\n", err.Error())
		return
	}

	backendService.GetLoadBalancer().Done(upstreamTarget)

	defer resp.Body.Close()

	for _, plugin := range g.plugs {
		if err := plugin.PostResponse(resp, g.reg, g.conf); err != nil {
			g.log.Infof("Plugin error: %v", err)
			http.Error(w, err.Error(), err.StatusCode())
			return
		}
	}

	// Log a message indicating that the response has been received from the target service
	g.log.Infof("Response received from %s: %d %s", upstreamTarget, resp.StatusCode, resp.Status)

	// Copy the response headers back to the client
	copyHeaders(w.Header(), resp.Header)

	// Set the status code and body of the response
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)

}

func copyHeaders(dst, src http.Header) {
	for k, v := range src {
		dst[k] = v
	}
}
