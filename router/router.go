package router

import (
	"go-admin/app/api"
	"go-admin/middleware"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.BindMiddlewareDefault(middleware.CORS)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.POST("/login", api.Auth.Login) // 登录
	})

	// auth
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.ALL("/hello", api.Hello)
		group.POST("/user", api.User.Create) // 创建一个用户
	})
}
