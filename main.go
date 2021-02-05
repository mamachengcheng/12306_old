package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/routers"
	"github.com/mamachengcheng/12306/app/static"
)

func main() {

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	routers.InitRouter(router)
	models.InitModel()

	//service.Start()

	router.Run(static.ServerAddress)
}
