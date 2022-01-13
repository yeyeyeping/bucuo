package main

import (
	"bucuo/middleware"
	"bucuo/routes"
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
	defer r.Run(":8080")
	AuthGroup("api/auth", r, routes.AuthAdminRoutesInit)
	UnAuthGroup("api", r, routes.UserRoutesInit)
}
