package app

import (
	"flybitch/app/controller"
	"flybitch/app/middleware"
	"flybitch/model"
)

func addRoutes() {
	api := e.Group("api")
	api.Use(middleware.Auth)
	api.GET("/ping", controller.Ping)
	api.POST("/print/query",controller.Query)
	api.POST("/print/body",controller.Analysis)
	api.POST("/SQL/adduser",model.Adduser)
	api.POST("/SQL/deteleuser",model.Deteleuser)
	api.POST("/SQL/find",model.Myfind)
	//api.POST("/SQL/modify",model.Modify)
}


