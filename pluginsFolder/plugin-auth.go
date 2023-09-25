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
	Client pb.AuthServiceClient
}

var GRPC_CLI *GRPCClient

type FErr struct {
	Err  string
	Code int
}

func (p *FErr) StatusCode() int {
	return p.Code
}

func (p *FErr) Error() string {
	return p.Err
}

func (p *FPlugin) PreRequest(req *http.Request, sr service.ServiceRegistry, cfg *config.Config) plugins.PluginError {
	if InArray(req.URL.Path, SkipRoutes) {
		return nil
	}
	if GRPC_CLI == nil {
		newGRPCClient()
	}
	response, err := GRPC_CLI.Client.VerifyToken(req.Context(), &pb.TokenRequest{Token: getToken(req.Header.Get("Authorization"))})
	if err != nil {
		return &FErr{
			Err:  err.Error(),
			Code: 500,
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
	conn, err := grpc.Dial("aaa-service:8001", grpc.WithInsecure())
	if err != nil {
		log.Error(err.Error())
	}
	GRPC_CLI = &GRPCClient{
		Client: pb.NewAuthServiceClient(conn),
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
