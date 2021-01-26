package serializers

type UserInformation struct {
	Name                string    `json:"name"`
	CertificateType     uint      `json:"certificate_type"`
	Sex                 bool      `json:"sex"`
	Birthday            string `json:"birthday"`
	Country             string    `json:"country"`
	CertificateDeadline string `json:"certificate_deadline"`
	Certificate         string    `json:"certificate"`
	PassengerType       uint      `json:"passenger_type"`
	MobilePhone         string    `json:"mobile_phone"`
	Email               string    `json:"email"`
	CheckStatus         uint      `json:"check_status"`
	UserStatus          uint      `json:"user_status"`
}

type User struct {
	UserName        string          `json:"user_name"`
	Password        string          `json:"password"`
	UserInformation UserInformation `json:"user_information"`
}

