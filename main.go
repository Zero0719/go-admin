package main

import (
	_ "go-admin/boot"
	_ "go-admin/router"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

func main() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	s.Run()
}
