package serializers

type Passenger struct {
	PassengerID uint `json:"passenger_id"`
}

type BookTickets struct {
	ScheduleID uint        `json:"schedule_id"`
	SeatType   uint        `json:"seat_type"`
	Passengers []Passenger `json:"passengers"`
}
