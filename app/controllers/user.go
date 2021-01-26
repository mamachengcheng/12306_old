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

func Register(c *gin.Context) {

	user := serializers.User{}
	c.BindJSON(&user)

	const format = "2006-01-02"
	birthday, _ := time.Parse(format, user.UserInformation.Birthday)
	certificateDeadline, _ := time.Parse(format, user.UserInformation.CertificateDeadline)

	utils.MysqlDB.Create(&models.User{UserName: user.UserName, Password: user.Password, UserInformation: models.Passenger{
		Name:                user.UserInformation.Name,
		CertificateType:     user.UserInformation.CertificateType,
		Sex:                 user.UserInformation.Sex,
		Birthday:            birthday,
		Country:             user.UserInformation.Country,
		CertificateDeadline: certificateDeadline,
		Certificate:         user.UserInformation.Certificate,
		PassengerType:       user.UserInformation.PassengerType,
		MobilePhone:         user.UserInformation.MobilePhone,
		Email:               user.UserInformation.Email,
		CheckStatus:         user.UserInformation.CheckStatus,
		UserStatus:          user.UserInformation.UserStatus,
	}})
}

func Login(c *gin.Context) {

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
	result := utils.MysqlDB.Preload("UserInformation").Where("user_name = ? AND password = ?", userName, password).Find(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.Code = 201
		response.Msg = "请正确输入用户名或密码"
	}

	token, _ := middlewares.GenerateToken(userName)
	response.Data.(map[string]interface{})["token"] = token

	utils.StatusOKResponse(response, c)
}
