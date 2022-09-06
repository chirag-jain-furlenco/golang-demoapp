package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	usersRoutes "demoapp/api/users"
	apiModel "demoapp/models/api"
)

func createRouterGroups(router *gin.RouterGroup) func(apiModel.SRouteGroupDef) {
	return func(routeGroupDef apiModel.SRouteGroupDef) {
		routeGroup := router.Group(routeGroupDef.Path)

		for _, route := range routeGroupDef.Routes {
			switch route.Method {
			case "GET":
				{
					routeGroup.GET(route.Route, route.Controller)
					break
				}
			case "POST":
				{
					routeGroup.POST(route.Route, route.Controller)
					break
				}
			case "PUT":
				{
					routeGroup.PUT(route.Route, route.Controller)
					break
				}
			case "DELETE":
				{
					routeGroup.DELETE(route.Route, route.Controller)
					break
				}
			default:
				{
					routeGroup.GET(route.Route, route.Controller)
					break
				}
			}
		}
	}
}

func StartServer() {
	server := gin.Default()
	port := viper.GetString("PORT")

	api := server.Group("/api")

	registerRouterGroups := createRouterGroups(api)

	registerRouterGroups(usersRoutes.Routes())

	server.Run(fmt.Sprintf(":%v", port))
}
