package static

const (
	IdCard = iota
	MainlandTravelPermitForHongKongAndMacaoResidents
	MainlandTravelPermitForTaiwanResidents
	Passport
	ForeignPermanentResidentIDCard
	HongKongMacaoAndTaiwanResidentResidencePermit
)

var CertificateType = []string{"中国居民身份证", "港澳居民来往内地通行证", "台湾居民来往大陆通行证", "护照", "外国人永久居留身份证", "港澳台居民居住证"}
