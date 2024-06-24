package controller

import (
	"apiserverv2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignupHandler(c *gin.Context) {
	/*
	   获取参数 -> 校验数据格式 -> 注册用户
	*/

	// 1.获取参数 2.校验数据格式
	var registerForm model.RegisterForm
	if err := c.ShouldBindJSON(&registerForm); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "error",
		})
	}

	// 3.注册用户
	

}
