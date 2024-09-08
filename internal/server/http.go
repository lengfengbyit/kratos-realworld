package server

import (
	"context"
	v1 "kratos-realworld/api/helloworld/v1"
	user "kratos-realworld/api/user/v1"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/middleware/auth"
	"kratos-realworld/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang-jwt/jwt/v5"
)

func MatchRouter() selector.MatchFunc {
	skipRouters := map[string]struct{}{
		"/user.v1.UserApi/Login":    {},
		"/user.v1.UserApi/Register": {},
	}
	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, userApi *service.UserApiService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(auth.Server(func(token *jwt.Token) (interface{}, error) {
				return []byte(c.Auth.Jwt.Secret), nil
			})).Match(MatchRouter()).Build(),
			validate.Validator(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	user.RegisterUserApiHTTPServer(srv, userApi)
	return srv
}
