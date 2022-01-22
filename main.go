package main

import (
	"bucuo/middleware"
	"bucuo/routes"
	"bucuo/util/setting"
	"github.com/gin-gonic/gin"
)

type RoutesFunc func(*gin.RouterGroup)

func AuthGroup(path string, route *gin.Engine, rfunc ...RoutesFunc) {
	auth := route.Group(path)
	auth.Use(middleware.JwtMiddleWare())
	{
		for _, routesFunc := range rfunc {
			routesFunc(auth)
		}
	}
}
func UnAuthGroup(path string, route *gin.Engine, rfunc ...RoutesFunc) {
	auth := route.Group(path)
	for _, routesFunc := range rfunc {
		routesFunc(auth)
	}
}
func main() {
	r := gin.Default()
	defer r.Run(setting.Port)
	r.Static("/static", "static")
	AuthGroup("api/auth", r,
		routes.AuthAdminRoutesInit,
		routes.AuthResouceRoutesInit,
		routes.AuthUserRoutesInit,
		routes.AuthExprRoutesInit,
		routes.AuthCommonRoutesInit,
	)
	UnAuthGroup("api", r,
		routes.UserRoutesInit,
		routes.ResouceRoutesInit,
		routes.ExprRoutesInit,
		routes.CommonRoutesInit,
	)
}
