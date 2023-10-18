package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/Frontman-Labs/frontman/config"
	"github.com/Frontman-Labs/frontman/log"
	"github.com/Frontman-Labs/frontman/pb"
	"github.com/Frontman-Labs/frontman/plugins"
	"github.com/Frontman-Labs/frontman/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type FPlugin struct{}

var FrontmanPlugin FPlugin

func (p *FPlugin) Name() string {
	return "FinS Auth Plugin"
}

var SkipRoutes = []string{
	"/health",
	"/aaa/v1/login",
	"/aaa/v1/token/refresh",
	"/aaa/v1/ui-config",
}

var FinSHeaders = []string{
	"X-Database-User",
	"X-Database-Password",
	"X-Database-Host",
	"X-Database-Port",
	"X-Database-Name",
}

type GRPCClient struct {
	Client pb.TokenServiceClient
}

var GRPC_CLI *GRPCClient

type FErr struct {
	Err  string
	Code int
	JSON any
}

func (p *FErr) StatusCode() int {
	return p.Code
}

func (p *FErr) Error() string {
	return p.Err
}

func (p *FErr) Response() any {
	return p.JSON
}

const ERR_TOKEN_IS_EMPTY = "token is empty"
const ERR_TOKEN_IS_INVALID = "token is invalid"
const ERR_TOKEN_IS_EXPIRED = "token is expired"

func (p *FPlugin) PreRequest(req *http.Request, sr service.ServiceRegistry, cfg *config.Config) plugins.PluginError {
	if InArray(req.URL.Path, SkipRoutes) {
		return nil
	}
	if GRPC_CLI == nil {
		newGRPCClient()
	}
	response, err := GRPC_CLI.Client.VerifyToken(req.Context(), &pb.VerifyTokenRequest{Token: getToken(req.Header.Get("Authorization"))})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if InArray(e.Message(), []string{ERR_TOKEN_IS_EMPTY, ERR_TOKEN_IS_INVALID, ERR_TOKEN_IS_EXPIRED}) {
				code := "ERR_TOKEN_IS_INVALID"
				if err.Error() == ERR_TOKEN_IS_EXPIRED {
					code = "ERR_TOKEN_IS_EXPIRED"
				}
				return &FErr{
					Err:  "unauthorize",
					Code: http.StatusUnauthorized,
					JSON: map[string]any{
						"message": "unauthorize",
						"code":    code,
					},
				}
			}
		}
		return &FErr{
			Err:  err.Error(),
			Code: http.StatusInternalServerError,
			JSON: map[string]any{
				"message": "unauthorize",
				"code":    "AAA_INTERNAL_SERVER_ERROR",
			},
		}
	}
	data := response.GetData()
	req.Header.Set("X-Database-User", data.DatabaseUser)
	req.Header.Set("X-Database-Password", data.DatabasePassword)
	req.Header.Set("X-Database-Host", data.DatabaseHost)
	req.Header.Set("X-Database-Port", fmt.Sprintf("%d", data.DatabasePort))
	req.Header.Set("X-Database-Name", data.DatabaseName)
	return nil
}

func (p *FPlugin) PostResponse(resp *http.Response, sr service.ServiceRegistry, cfg *config.Config) plugins.PluginError {
	for _, v := range FinSHeaders {
		resp.Header.Del(v)
	}
	return nil
}

func (p *FPlugin) Close() plugins.PluginError {
	// Cleanup resources used by the plugin
	return nil
}

func newGRPCClient() {
	conn, err := grpc.Dial("aaa-service:8000", grpc.WithInsecure())
	if err != nil {
		log.Error(err.Error())
	}
	GRPC_CLI = &GRPCClient{
		Client: pb.NewTokenServiceClient(conn),
	}
}

func getToken(authorizationHeader string) string {
	return strings.Replace(authorizationHeader, "Bearer ", "", 1)
}

func InArray(item any, array any) bool {
	arr := reflect.ValueOf(array)
	if arr.Kind() != reflect.Slice {
		return false
	}
	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}
