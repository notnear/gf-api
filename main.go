package main

import (
	_ "gf-api/boot"
	"gf-api/common"
	_ "gf-api/router"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/swagger"
)

// @title       `gf-api` 基于 GF(Go Frame) api demo
// @version     1.0
// @description `gf-api` api demo 接口文档。
// @schemes     http https
func main() {
	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	server := g.Server()
	server.BindStatusHandlerByMap(map[int]ghttp.HandlerFunc{
		403: func(r *ghttp.Request) { common.JsonExit(r, common.Forbid) },
		404: func(r *ghttp.Request) { common.JsonExit(r, common.NoRoute) },
		500: func(r *ghttp.Request) { common.JsonExit(r, common.SystemError) },
	})

	server.Plugin(&swagger.Swagger{})

	server.Run()
}
