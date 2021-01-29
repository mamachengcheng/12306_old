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
