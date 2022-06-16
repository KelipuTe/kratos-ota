package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

var currentUserKey struct{}

type CurrentUser struct {
	Username string
}

func WithContext(ctx context.Context, user *CurrentUser) context.Context {
	return context.WithValue(ctx, currentUserKey, user)
}

// MakeJWT 构造 JWT
func MakeJWT(secret string, username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return tokenStr
}

func JWTAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// 从请求头获取 token
				tokenStr := tr.RequestHeader().Get("Authorization")
				// 拆分 token 得到 JWT
				arr1Auth := strings.SplitN(tokenStr, " ", 2)
				if len(arr1Auth) != 2 || !strings.EqualFold(arr1Auth[0], "Token") {
					return nil, errors.New("jwt token missing")
				}
				// 解析 JWT
				token, err := jwt.Parse(arr1Auth[1], func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(secret), nil
				})
				if err != nil {
					return nil, err
				}
				// 从 JWT 拿数据
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					spew.Dump(claims)
					// 生成一个新的 Context 把数据放进去
					if u, ok := claims["username"]; ok {
						ctx = WithContext(ctx, &CurrentUser{Username: u.(string)})
					}
				} else {
					return nil, errors.New("Token Invalid")
				}

				defer func() {
					// Do something on exiting
				}()
			}
			return handler(ctx, req)
		}
	}
}
