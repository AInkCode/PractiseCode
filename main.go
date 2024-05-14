package main

import (
	"PratiseCode/dao"
	"PratiseCode/model"
	"PratiseCode/routers"
	"PratiseCode/setting"
	"fmt"
)

/*
	url     --> controller  --> logic   -->    model

请求来了  -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/

const confFile = "E:\\CodeHub\\GoHub\\src\\PratiseCode\\config.ini"

func main() {
	if err := setting.Init(confFile); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	err := dao.InitMysql(setting.Config.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close() // 程序退出关闭数据库连接
	// 模型绑定
	dao.DB.AutoMigrate(&model.Todo{})
	r := routers.SetupRouters()
	r.Run(":80")
}
