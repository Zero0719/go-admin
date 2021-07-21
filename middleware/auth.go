package middleware

import (
	"go-admin/package/jwt"
	"go-admin/package/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

func Auth(r *ghttp.Request) {
	authorization := r.Header.Get("Authorization")

	if authorization == "" {
		response.UnAuthorized(r, "缺少token")
	}

	token := gstr.Split(authorization, " ")
	tokenString := token[1]

	_, err := jwt.ParseToken(tokenString)

	if err != nil {
		response.UnAuthorized(r, err.Error())
	}

	r.Middleware.Next()
}
