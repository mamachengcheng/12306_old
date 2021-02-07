package serializers

type Register struct {
	Username    string `validate:"required,VerifyUsernameFormat" json:"username"`
	Password    string `validate:"required,VerifyPasswordFormat" json:"password"`
	Email       string `validate:"required,email" json:"email"`
	MobilePhone string `validate:"required,numeric,len=11,VerifyMobilePhoneFormat" json:"mobile_phone"`
	Name        string `validate:"required" json:"name"`
	Certificate string `validate:"required" json:"certificate"`
}

type Login struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required,VerifyPasswordFormat" json:"password"`
}

type QueryUserInformation struct {
	Username        string `json:"username"`
	Name            string `json:"name"`
	Country         string `json:"country"`
	CertificateType string `json:"certificate_type"`
	Certificate     string `json:"certificate"`
	CheckStatus     string `json:"check_status"`
	MobilePhone     string `json:"mobile_phone"`
	Email           string `json:"email"`
	PassengerType   string `json:"passenger_type"`
}

type QueryRegularPassenger struct {
	CertificateType string `json:"certificate_type"`
	Name            string `json:"name"`
	Certificate     string `json:"certificate"`
	PassengerType   string `json:"passenger_type"`
	CheckStatus     string `json:"check_status"`
	CreateDate      string `json:"create_date"`
	MobilePhone     string `json:"mobile_phone"`
}

type AddRegularPassenger struct {
	MobilePhone string `validate:"required,numeric,len=11,VerifyMobilePhoneFormat" json:"mobile_phone"`
	Name        string `validate:"required,VerifyNameFormat" json:"name"`
	Certificate string `validate:"required,len=18,VerifyCertificateFormat" json:"certificate"`
}

type DeleteRegularPassenger struct {
	PassengerID uint `validate:"required" json:"passenger_id"`
}

type UpdateRegularPassenger struct {
	PassengerID   uint   `validate:"required" json:"passenger_id"`
	MobilePhone   string `validate:"required,numeric,len=11,VerifyMobilePhoneFormat" json:"mobile_phone"`
	PassengerType uint8  `validate:"required" json:"passenger_type"`
}

type UpdatePassword struct {
	Password string `validate:"required,VerifyPasswordFormat" json:"password""`
}
