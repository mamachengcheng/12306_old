package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/resource"
	"github.com/mamachengcheng/12306/app/routers"
	"github.com/mamachengcheng/12306/app/service"
	"gopkg.in/ini.v1"
)

func main() {
	gin.SetMode(gin.DebugMode)


	cfg, _ := ini.Load(resource.ConfFilePath)
	server := cfg.Section("server")

	router := gin.Default()

	routers.InitRouter(router)
	models.InitModel()
	service.Start()

	address := server.Key("http").String() + ":" + server.Key("port").String()
	router.Run(address)
}
