package main

import "github.com/mamachengcheng/12306/app/service"

func main() {
	//gin.SetMode(gin.DebugMode)
	//router := gin.Default()
	//routers.InitRouter(router)
	//router.Run(static.ServerAddress)

	service.Start()
}
