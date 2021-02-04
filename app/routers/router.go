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
		user.POST("/update_password", controllers.UpdatePasswordAPI)

		user.POST("/add_regular_passengers", controllers.AddRegularPassengerAPI)
		user.GET("/query_regular_passengers", controllers.QueryRegularPassengersAPI)
		user.POST("/update_regular_passenger", controllers.UpdateRegularPassengerAPI)
		user.POST("/delete_regular_passenger", controllers.DeleteRegularPassengerAPI)
	}

	// Train part router.
	train := router.Group("/api/v1/train")
	train.GET("get_station_list", controllers.GetStationListAPI)
	train.GET("search_station", controllers.SearchStationAPI)
	train.GET("get_schedule_list", controllers.GetScheduleListAPI)
	train.GET("get_stops", controllers.GetStopAPI)

	// Order part router.
	order := router.Group("/api/v1/user")
	order.Use(middlewares.JWTMiddleware())
	{
		order.GET("refund_money", controllers.ReadyPayAPI)

		order.POST("book_ticket", controllers.BookTicketAPI)
		order.POST("cancel_order", controllers.CancelOrderAPI)
		order.POST("pay_order", controllers.PayOrderAPI)
		order.POST("refund_ticket", controllers.RefundTicketAPI)
	}
}
