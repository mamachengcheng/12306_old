package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/middlewares"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/resource"
	"github.com/mamachengcheng/12306/app/serializers"
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
)

func RegisterAPI(c *gin.Context) {

	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "注册成功",
	}

	data := serializers.Register{}
	c.BindJSON(&data)

	// 输入合法性检验
	validate := serializers.GetValidate()
	err := validate.Struct(data)

	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		// 用户存在检验
		user := models.User{}
		err = utils.MysqlDB.Where("username = ?", data.Username).First(&user).Error

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			response.Code = 202
			response.Msg = "用户已存在"
		} else {

			// 进行注册
			sex, birthday := utils.ParseIdentityCard(data.Certificate)

			utils.MysqlDB.Create(&models.User{
				Username:    data.Username,
				Email:       data.Email,
				MobilePhone: data.MobilePhone,
				Password:    data.Password,

				UserInformation: models.Passenger{
					Name:        data.Name,
					Sex:         sex,
					Birthday:    birthday,
					Certificate: data.Certificate,
					MobilePhone: data.MobilePhone,
				}})
		}
	}

	utils.StatusOKResponse(response, c)
}

func LoginAPI(c *gin.Context) {

	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "登陆成功",
	}

	data := serializers.Login{}
	c.BindJSON(&data)

	// 输入合法性检验
	validate := serializers.GetValidate()
	err := validate.Struct(data)

	if err != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {

		// 用户存在检验
		user := models.User{}
		err := utils.MysqlDB.Preload("QueryUserInformation").Where("username = ? OR email = ? OR mobile_phone = ?", data.Username, data.Username, data.Username).Where("password = ?", data.Password).First(&user).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Code = 202
			response.Msg = "请正确输入用户名或密码"
		} else {
			token, _ := middlewares.GenerateToken(user.Username)
			response.Data.(map[string]interface{})["token"] = token
			response.Data.(map[string]interface{})["user_information"] = user.UserInformation
		}
	}

	utils.StatusOKResponse(response, c)
}

func QueryUserInformationAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "查询成功",
	}

	claims := c.MustGet("claims").(*middlewares.Claims)

	user := models.User{}
	utils.MysqlDB.Preload("QueryUserInformation").Where("username = ?", claims.Username).First(&user)

	response.Data.(map[string]interface{})["user_information"] = serializers.QueryUserInformation{
		Username:        user.Username,
		Name:            user.UserInformation.Name,
		Country:         user.UserInformation.Country,
		CertificateType: resource.CertificateType[user.UserInformation.CertificateType],
		Certificate:     user.UserInformation.Certificate,
		CheckStatus:     resource.CheckStatus[user.UserInformation.CheckStatus],
		MobilePhone:     user.MobilePhone,
		Email:           user.Email,
		PassengerType:   resource.PassengerType[user.UserInformation.PassengerType],
	}

	utils.StatusOKResponse(response, c)
}

func QueryRegularPassengersAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "查询成功",
	}

	claims := c.MustGet("claims").(*middlewares.Claims)

	user := models.User{}
	utils.MysqlDB.Preload("QueryUserInformation").Where("user_name = ?", claims.Username).First(&user)

	var passengers []serializers.QueryRegularPassenger

	for _, passenger := range user.Passengers {
		passengers = append(passengers, serializers.QueryRegularPassenger{
			CertificateType: resource.CertificateType[passenger.CertificateType],
			Name:            passenger.Name,
			Certificate:     passenger.Certificate,
			PassengerType:   resource.PassengerType[passenger.PassengerType],
			CheckStatus:     resource.PassengerType[passenger.CheckStatus],
			CreateDate:      passenger.CreatedAt.Format("2006-01-02"),
			MobilePhone:     passenger.MobilePhone,
		})
	}

	response.Data.(map[string]interface{})["passengers"] = passengers

	utils.StatusOKResponse(response, c)
}

func AddRegularPassengerAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "添加成功",
	}

	claims := c.MustGet("claims").(*middlewares.Claims)

	data := serializers.AddRegularPassenger{}
	c.BindJSON(&data)

	sex, birthday := utils.ParseIdentityCard(data.Certificate)

	user := models.User{}
	utils.MysqlDB.Where("username = ?", claims.Username).First(&user)
	utils.MysqlDB.Model(&user).Association("Passengers").Append(&models.Passenger{
		Name:        data.Name,
		Sex:         sex,
		Birthday:    birthday,
		Certificate: data.Certificate,
		MobilePhone: data.MobilePhone,
	})

	utils.StatusOKResponse(response, c)
}


// TODO: Complete UpdateRegularPassengerAPI.
func UpdateRegularPassengerAPI(c *gin.Context) {

}

// TODO: Complete DeleteRegularPassengerAPI.
func DeleteRegularPassengerAPI(c *gin.Context) {

}


// TODO: Complete UpdatePasswordAPI.
func UpdatePasswordAPI(c *gin.Context) {

}
