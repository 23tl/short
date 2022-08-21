// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	auth "go-zero-short/internal/handler/auth"
	testing "go-zero-short/internal/handler/testing"
	uri "go-zero-short/internal/handler/uri"
	user "go-zero-short/internal/handler/user"
	"go-zero-short/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/testing",
				Handler: testing.TestingHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: auth.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: auth.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/auth"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/current",
				Handler: user.CurrentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/send",
				Handler: uri.SendHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/uri"),
	)
}
