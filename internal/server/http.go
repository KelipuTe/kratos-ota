package server

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	v1 "kratos-ota/api/ota/v1"
	"kratos-ota/internal/conf"
	internalError "kratos-ota/internal/error"
	"kratos-ota/internal/pkg/middleware/auth"
	"kratos-ota/internal/service"
	netHttp "net/http"
)

// NewSkipRouterMatcher 过滤不需要 JWT 鉴权的路由
func NewSkipRouterMatcher() selector.MatchFunc {
	mapSkipRouter := map[string]struct{}{
		"/ota.v1.Ota/GetToken": {},
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := mapSkipRouter[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, cauth *conf.Auth, ota *service.OtaService, logger log.Logger) *http.Server {
	http.ErrorEncoder(t1ErrorEncoder)

	var opts = []http.ServerOption{
		http.Filter(
			// 过滤器
			func(h netHttp.Handler) netHttp.Handler {
				return netHttp.HandlerFunc(func(writer netHttp.ResponseWriter, request *netHttp.Request) {
					spew.Dump("in")
					h.ServeHTTP(writer, request)
					spew.Dump("out")
				})
			},
			// CORS 跨域
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"baidu"}),
			),
		),
		http.Middleware(
			recovery.Recovery(),
			// JWT 鉴权
			selector.Server(auth.JWTAuth(cauth.Jwt.SecretKey)).Match(NewSkipRouterMatcher()).Build(),
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
	v1.RegisterOtaHTTPServer(srv, ota)
	return srv
}

func t1ErrorEncoder(w netHttp.ResponseWriter, r *netHttp.Request, err error) {
	t1err := internalError.FromHTTPError(err)

	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(t1err)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/"+codec.Name())
	if t1err.Code > 99 && t1err.Code < 600 {
		w.WriteHeader(int(t1err.Code))
	} else {
		w.WriteHeader(500)
	}

	_, _ = w.Write(body)
}
