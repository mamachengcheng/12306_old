package serializers

import "regexp"
import "github.com/go-playground/validator/v10"

func GetValidate() *validator.Validate {

	Validate := validator.New()
	Validate.RegisterValidation("VerifyMobilePhoneFormat", VerifyMobilePhoneFormat)
	Validate.RegisterValidation("VerifyDateFormat", VerifyDateFormat)
	Validate.RegisterValidation("VerifyUsernameFormat", VerifyUsernameFormat)
	Validate.RegisterValidation("VerifyPasswordFormat", VerifyPasswordFormat)
	Validate.RegisterValidation("VerifyCertificateFormat", VerifyCertificateFormat)
	Validate.RegisterValidation("VerifyNameFormat", VerifyNameFormat)

	Validate.RegisterValidation("VerifyInitialNameFormat", VerifyInitialNameFormat)

	Validate.RegisterValidation("VerifyTrainNoFormat", VerifyTrainNoFormat)
	Validate.RegisterValidation("VerifyTimeFormat", VerifyTimeFormat)
	Validate.RegisterValidation("VerifyStationNameFormat", VerifyStationNameFormat)
	Validate.RegisterValidation("VerifyTrainNoFormat", VerifyTrainNoFormat)
	return Validate
}

func VerifyMobilePhoneFormat(mobilePhone validator.FieldLevel) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobilePhone.Field().String())
}

func VerifyDateFormat(date validator.FieldLevel) bool {
	regular := "^\\d{4}-\\d{2}-\\d{2}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(date.Field().String())
}

func VerifyUsernameFormat(username validator.FieldLevel) bool {
	regular := "^\\w{6,30}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(username.Field().String())
}

func VerifyPasswordFormat(password validator.FieldLevel) bool {
	regular := "^[a-zA-Z0-9_]{6,20}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(password.Field().String())
}

func VerifyCertificateFormat(password validator.FieldLevel) bool {
	regular := "^\\d{6}(\\d{8})\\d{2}(\\d)[0-9X]$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(password.Field().String())
}

func VerifyNameFormat(password validator.FieldLevel) bool {
	regular := "^[\u4e00-\u9fa5]{2,6}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(password.Field().String())
}

func VerifyInitialNameFormat(initialName validator.FieldLevel) bool {
	regular := "[a-zA-Z]"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(initialName.Field().String())
}

func VerifyTimeFormat(date validator.FieldLevel) bool {
	regular := "^(([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})-(((0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8]))))$|^((([0-9]{2})(0[48]|[2468][048]|[13579][26])|((0[48]|[2468][048]|[13579][26])00))-02-29)$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(date.Field().String())
}

func VerifyStationNameFormat(stationName validator.FieldLevel) bool {
	regular := "^[\u4e00-\u9fa5]{2,12}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(stationName.Field().String())
}

func VerifyTrainNoFormat(trainNo validator.FieldLevel) bool {
	regular := "^[DGKTZ0-9]{2,6}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(trainNo.Field().String())
}
