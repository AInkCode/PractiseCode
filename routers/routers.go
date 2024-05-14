package routers

import (
	"PratiseCode/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()
	// v1
	v1Group := r.Group("v1")
	{
		// 功能区

		// 增
		v1Group.POST("/todo", controller.AddTodo)
		// 删
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
		// 改
		v1Group.PUT("/todo/:id", controller.UpdataTodo)
		// 查
		v1Group.GET("/todo/:id", controller.GetTodo)
	}
	return r
}
