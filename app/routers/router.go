package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/controllers"
	"github.com/mamachengcheng/12306/app/middlewares"
)

func InitRouter(router *gin.Engine) {
	// User part router.
	user := router.Group("/api/v1/user")
	user.POST("/register", controllers.RegisterAPI)
	user.POST("/login", controllers.LoginAPI)
	user.Use(middlewares.JWTMiddleware())
	{
		user.GET("/query_user_information", controllers.QueryUserInformationAPI)
		user.GET("/query_regular_passengers", controllers.QueryRegularPassengersAPI)
		user.POST("/add_regular_passengers", controllers.AddRegularPassengersAPI)
	}

	// Train part router.
	//train := router.Group("/api/v1/user")

	// Order part router.
	order := router.Group("/api/v1/user")
	order.Use(middlewares.JWTMiddleware())
	{

	}
}
