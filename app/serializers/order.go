package serializers

type BookTicket struct {
	ScheduleID  uint `json:"schedule_id"`
	SeatType    uint `json:"seat_type"`
	PassengerID uint `json:"passenger_id"`

	SeatPrefer  uint `json:"seat_prefer"`
}
