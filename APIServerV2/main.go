package main

import "apiserverv2/routes"

/*
  API -> routes -> model -> mysql -> redis <- dao
*/

func main() {
	// 创建路由
	routes := routes.SetupRoutes()
}
