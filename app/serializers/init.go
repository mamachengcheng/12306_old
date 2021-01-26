package serializers

import "regexp"
import "github.com/go-playground/validator/v10"

func init() {
	Validate := validator.New()
	Validate.RegisterValidation("VerifyMobilePhoneFormat", VerifyMobilePhoneFormat)
	Validate.RegisterValidation("VerifyDateFormat", VerifyDateFormat)
	Validate.RegisterValidation("VerifyUsernameFormat", VerifyUsernameFormat)
	Validate.RegisterValidation("VerifyPasswordFormat", VerifyPasswordFormat)
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
	regular := "^\\d{4}-\\d{2}-\\d{2}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(username.Field().String())
}

func VerifyPasswordFormat(password validator.FieldLevel) bool {
	regular := "^\\d{4}-\\d{2}-\\d{2}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(password.Field().String())
}
