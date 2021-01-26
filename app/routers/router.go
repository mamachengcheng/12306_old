package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/controllers"
	"github.com/mamachengcheng/12306/app/middlewares"
)

func InitRouter(router *gin.Engine) {
	// User part router.
	user := router.Group("/api/v1/user")
	user.POST("/register", controllers.Register)
	user.POST("/login", controllers.Login)
	user.Use(middlewares.JWTMiddleware())
	{

	}

	// Train part router.
	//train := router.Group("/api/v1/user")

	// Order part router.
	order := router.Group("/api/v1/user")
	order.Use(middlewares.JWTMiddleware())
	{

	}
}
