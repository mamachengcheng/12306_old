package serializers

type UserInformation struct {
	Name                string `validate:"required" json:"name"`
	CertificateType     uint8  `validate:"required,gte=0,lte=3" json:"certificate_type"`
	Sex                 bool   `validate:"required" json:"sex"`
	Birthday            string `validate:"required,VerifyDateFormat" json:"birthday"`
	Country             string `validate:"required" json:"country"`
	CertificateDeadline string `validate:"required,VerifyDateFormat" json:"certificate_deadline"`
	Certificate         string `validate:"required" json:"certificate"`
	PassengerType       uint8  `validate:"required,gte=0,lte=3" json:"passenger_type"`
	MobilePhone         string `validate:"required,numeric,len=11,VerifyMobilePhoneFormat" json:"mobile_phone"`
}

type User struct {
	Username        string          `validate:"required,VerifyUsernameFormat" json:"user_name"`
	Password        string          `validate:"required,VerifyPasswordFormat" json:"password"`
	Mail            string          `validate:"required,email" json:"mail"`
	UserInformation UserInformation `validate:"required" json:"user_information"`
}

type Passenger struct {
	Name                string `validate:"required" gorm:"not null" json:"name"`
	CertificateType     uint   `validate:"required,gte=0,lte=3" json:"certificate_type"`
	Sex                 bool   `validate:"required" json:"sex"`
	Birthday            string `validate:"required,VerifyDateFormat" json:"birthday"`
	Country             string `validate:"required" json:"country"`
	CertificateDeadline string `validate:"required,VerifyDateFormat" json:"certificate_deadline"`
	Certificate         string `validate:"required" json:"certificate"`
	PassengerType       uint   `validate:"required,gte=0,lte=3" json:"passenger_type"`
	MobilePhone         string `validate:"required,numeric,len=11,VerifyMobilePhoneFormat" json:"mobile_phone"`
}
