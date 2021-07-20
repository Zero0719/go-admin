package main

import (
	_ "go-admin/boot"
	_ "go-admin/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
