package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/resource"
	"github.com/mamachengcheng/12306/app/serializers"
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
	"log"
)

func Register(c *gin.Context) {
	//var (
	//	registerSuccess = utils.SubStatus{Code: "register-success", Msg: "注册成功"}
	//	//RegisteredError = utils.SubStatus{Code: "registered-error", Msg: "已注册"}
	//)

	data := make(map[string]interface{})
	c.BindJSON(&data)
	log.Printf("%v",&data)

	//user := &models.User{}
	//
	//if utils.MysqlDBErr == nil {
	//	utils.MysqlDB.Create(&user)
	//	utils.DefaultResponse(resource.Success, registerSuccess.Code, nil, registerSuccess.Msg, c)
	//} else {
	//	utils.DefaultResponse(resource.Success, registerSuccess.Code, nil, registerSuccess.Msg, c)
	//}
}

func Login(c *gin.Context) {
	var (
		loginSuccess        = utils.SubStatus{Code: "login-success", Msg: "登陆成功"}
		parameterError      = utils.SubStatus{Code: "parameter-error", Msg: "请正确输入用户名或密码"}
		incorrectInputError = utils.SubStatus{Code: "incorrect-input-error", Msg: "请正确输入用户名或密码"}
	)

	var (
		LoginData serializers.Login
		user      models.User
	)

	if err := c.ShouldBind(&LoginData); err != nil {
		result := utils.MysqlDB.Where("username = ? AND password = ?", LoginData.Username, LoginData.Password).First(&user)
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.DefaultResponse(resource.Success, loginSuccess.Code, nil, loginSuccess.Msg, c)
		} else {
			utils.DefaultResponse(resource.ParameterError, incorrectInputError.Code, nil, incorrectInputError.Msg, c)
		}
	} else {
		utils.DefaultResponse(resource.ParameterError, parameterError.Code, nil, parameterError.Msg, c)
	}

}

