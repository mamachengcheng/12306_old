package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/middlewares"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/serializers"
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
	"time"
)

func RegisterAPI(c *gin.Context) {

	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "注册成功",
	}

	data := serializers.User{}
	c.BindJSON(&data)

	user := models.User{}
	result := utils.MysqlDB.Where("user_name = ?", data.UserName).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		const format = "2006-01-02"
		birthday, _ := time.Parse(format, data.UserInformation.Birthday)
		certificateDeadline, _ := time.Parse(format, data.UserInformation.CertificateDeadline)

		utils.MysqlDB.Create(&models.User{UserName: data.UserName, Password: data.Password, UserInformation: models.Passenger{
			Name:                data.UserInformation.Name,
			CertificateType:     data.UserInformation.CertificateType,
			Sex:                 data.UserInformation.Sex,
			Birthday:            birthday,
			Country:             data.UserInformation.Country,
			CertificateDeadline: certificateDeadline,
			Certificate:         data.UserInformation.Certificate,
			PassengerType:       data.UserInformation.PassengerType,
			MobilePhone:         data.UserInformation.MobilePhone,
			Email:               data.UserInformation.Email,
			CheckStatus:         data.UserInformation.CheckStatus,
			UserStatus:          data.UserInformation.UserStatus,
		}})
	} else {
		response.Code = 201
		response.Msg = "用户已存在"
	}

	utils.StatusOKResponse(response, c)
}

func LoginAPI(c *gin.Context) {

	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "登陆成功",
	}

	data := serializers.User{}
	c.BindJSON(&data)
	userName := data.UserName
	password := data.Password

	user := models.User{}
	result := utils.MysqlDB.Preload("UserInformation").Where("user_name = ? AND password = ?", userName, password).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.Code = 201
		response.Msg = "请正确输入用户名或密码"
	} else {
		token, _ := middlewares.GenerateToken(userName)
		response.Data.(map[string]interface{})["token"] = token
	}

	utils.StatusOKResponse(response, c)
}
