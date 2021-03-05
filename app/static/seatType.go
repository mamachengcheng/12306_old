package static

const (
	HardSeat = iota
	SoftSeat
	HardSleeper
	SoftSleeper
	BusinessClass
	FirstClass
	SecondClass
	NoSeat
)

var SeatType = []string{"硬座", "软座", "硬卧", "软卧", "商务座", "一等座", "二等座", "无座"}
