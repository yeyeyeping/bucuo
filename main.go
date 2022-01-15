package main

import (
	"bucuo/middleware"
	"bucuo/routes"
	"bucuo/util"
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
	gin.ForceConsoleColor()
	r := gin.Default()
	r.Static("/static", "static")
	defer r.Run(util.Port)
	AuthGroup("api/auth", r, routes.AuthAdminRoutesInit)
	UnAuthGroup("api", r, routes.UserRoutesInit)
	//r := model.Reply{
	//	Content: "12313113",
	//}
	//dao.DB.Create(&r)
	//fmt.Printf("%#v", r)
}
