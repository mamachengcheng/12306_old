package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mamachengcheng/12306/app/middlewares"
	"github.com/mamachengcheng/12306/app/utils"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func BookTicketAPI(c *gin.Context) {
	claims := c.MustGet("claims").(*middlewares.Claims)
	if err := utils.RedisDB.Get(utils.RedisDBCtx, claims.Username).Err(); err == nil {
		utils.RedisDB.Set(utils.RedisDBCtx, claims.Username, true, time.Minute*30)
	}

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)

	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "",
	}

	if err == nil {
		defer ws.Close()
		for {
			response.Data.(map[string]interface{})["remaining_time"] = utils.RedisDB.TTL(utils.RedisDBCtx, claims.Username)
			_ = ws.WriteJSON(response)
		}
	}
}

func RefundTicketAPI(c *gin.Context) {

}

func PayOrderAPI(c *gin.Context) {

}

func RefundMoneyAPI(c *gin.Context) {

}
