package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/middlewares"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/serializers"
	"github.com/mamachengcheng/12306/app/static"
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
			user := models.User{
				Username:    data.Username,
				Email:       data.Email,
				MobilePhone: data.MobilePhone,
				Password:    data.Password,
			}
			passenger := models.Passenger{
				Name:        data.Name,
				Sex:         sex,
				Birthday:    birthday,
				Certificate: data.Certificate,
				MobilePhone: data.MobilePhone,
			}

			utils.MysqlDB.Create(&user)

			if err := utils.MysqlDB.Model(&user).Association("Passengers").Append(&passenger); err == nil {
				utils.MysqlDB.Model(&user).Update("user_information_id", passenger.ID)
			}
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
		err := utils.MysqlDB.Where("username = ? OR email = ? OR mobile_phone = ?", data.Username, data.Username, data.Username).Where("password = ?", data.Password).First(&user).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Code = 202
			response.Msg = "请正确输入用户名或密码"
		} else {
			token, _ := middlewares.GenerateToken(user.Username)
			response.Data.(map[string]interface{})["token"] = token
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

	var user models.User
	var passenger models.Passenger
	utils.MysqlDB.Where("username = ?", claims.Username).First(&user)
	utils.MysqlDB.Where("id = ?", user.UserInformationID).First(&passenger)

	response.Data.(map[string]interface{})["user_information"] = serializers.QueryUserInformation{
		Username:        user.Username,
		Name:            passenger.Name,
		Country:         passenger.Country,
		CertificateType: static.CertificateType[passenger.CertificateType],
		Certificate:     passenger.Certificate,
		CheckStatus:     static.CheckStatus[passenger.CheckStatus],
		MobilePhone:     user.MobilePhone,
		Email:           user.Email,
		PassengerType:   static.PassengerType[passenger.PassengerType],
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
	utils.MysqlDB.Preload("Passengers").Where("username = ?", claims.Username).First(&user)

	var passengers []serializers.QueryRegularPassenger

	for _, passenger := range user.Passengers {
		passengers = append(passengers, serializers.QueryRegularPassenger{
			CertificateType: static.CertificateType[passenger.CertificateType],
			Name:            passenger.Name,
			Certificate:     passenger.Certificate,
			PassengerType:   static.PassengerType[passenger.PassengerType],
			CheckStatus:     static.CheckStatus[passenger.CheckStatus],
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

func DeleteRegularPassengerAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "删除成功",
	}

	data := serializers.DeleteRegularPassenger{}
	c.BindJSON(&data)

	// 输入合法性检验
	if validate := serializers.GetValidate(); validate.Struct(data) != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		claims := c.MustGet("claims").(*middlewares.Claims)

		user := models.User{}
		utils.MysqlDB.Preload("Passengers", "id = ?", data.PassengerID).Where("username = ?", claims.Username).First(&user)

		if len(user.Passengers) == 0 {
			response.Code = 202
			response.Msg = "乘客不存在"
		} else {
			utils.MysqlDB.Unscoped().Where("ID = ?", data.PassengerID).Delete(&models.Passenger{})
		}
	}

	utils.StatusOKResponse(response, c)
}

func UpdateRegularPassengerAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "删除成功",
	}

	data := serializers.UpdateRegularPassenger{}
	c.BindJSON(&data)

	// 输入合法性检验
	if validate := serializers.GetValidate(); validate.Struct(data) != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		claims := c.MustGet("claims").(*middlewares.Claims)

		var user models.User
		var passenger models.Passenger
		utils.MysqlDB.Where("username = ?", claims.Username).First(&user)
		utils.MysqlDB.Where("id = ? AND user_refer = ?", data.PassengerID, user.ID).First(&passenger)
		utils.MysqlDB.Model(&passenger).Updates(map[string]interface{}{"mobile_phone": data.MobilePhone, "passenger_type": data.PassengerType})
	}

	utils.StatusOKResponse(response, c)
}

func UpdatePasswordAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "修改成功",
	}

	data := serializers.UpdatePassword{}
	c.BindJSON(&data)

	// 输入合法性检验
	if validate := serializers.GetValidate(); validate.Struct(data) != nil {
		response.Code = 201
		response.Msg = "参数不合法"
	} else {
		claims := c.MustGet("claims").(*middlewares.Claims)
		utils.MysqlDB.Model(&models.User{}).Where("username= ?", claims.Username).Update("password", data.Password)
	}

	utils.StatusOKResponse(response, c)
}
