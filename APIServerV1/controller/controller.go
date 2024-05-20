package controller

import (
	"PratiseCode/APIServerV1/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 增
func AddTodo(c *gin.Context) {
	// 1.获取链接数据
	var todo model.Todo
	// c.BindJSON(&todo)
	c.ShouldBindJSON(&todo)
	// 2.存储数据到数据库
	if err := model.AddATodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// 删
func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "request error",
		})
	} else {
		if err := model.DeleteATodo(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "delete error",
			})
		} else {
			c.JSON(http.StatusOK, id)
		}
	}
}

// 改
func UpdataTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id request",
		})
	}
	todo, err := model.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "get error",
		})
	}
	c.BindJSON(&todo)
	if err := model.UpdataATodo(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "updata error",
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// 查
func GetTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	fmt.Println(id)
	// id := c.Param("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id request",
		})
	}
	todo, err := model.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "get error",
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
	// todoList, err := model.GetAllTodo()
	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	// } else {
	// 	c.JSON(http.StatusOK, todoList)
	// }
}
