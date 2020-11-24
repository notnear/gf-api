package router

import (
	"gf-api/app/controller"
	"github.com/gogf/gf/frame/g"
)

func init() {
	s := g.Server()

	userGroup := s.Group("/user")
	userGroup.ALL("/add", controller.ActionAdd)
	userGroup.ALL("/list", controller.ActionUserList)
}
