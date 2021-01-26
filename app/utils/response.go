package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"` // 业务返回码
	Data interface{} `json:"data"` //业务返回数据
	Msg  string      `json:"msg"`  // 业务返回描述
}

func StatusOKResponse(response Response, c *gin.Context) {
	c.JSON(http.StatusOK, response)
}
