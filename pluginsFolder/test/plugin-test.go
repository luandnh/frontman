package main

import (
	"net/http"
	_ "plugin"

	"github.com/Frontman-Labs/frontman/config"
	"github.com/Frontman-Labs/frontman/plugins"
	"github.com/Frontman-Labs/frontman/service"
)

type FPlugin struct{}

var FrontmanPlugin FPlugin

func (p *FPlugin) Name() string {
	return "FPlugin"
}

func (p *FPlugin) PreRequest(req *http.Request, sr service.ServiceRegistry, cfg *config.Config) plugins.PluginError {
	// Modify the request before sending it to the target service
	return nil
}

func (p *FPlugin) PostResponse(resp *http.Response, sr service.ServiceRegistry, cfg *config.Config) plugins.PluginError {
	// Modify the response before sending it back to the client
	return nil
}

func (p *FPlugin) Close() plugins.PluginError {
	// Cleanup resources used by the plugin
	return nil
}
