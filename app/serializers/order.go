package serializers

type CreateOrder struct {
	ScheduleID uint64   `json:"schedule_id"`
	SeatType   uint64   `json:"seat_type"`
	Passengers []uint64 `json:"passengers"`
}

type CancelOrder struct {
	OrderID uint64 `json:"order_id"`
}

type PayMoney struct {
	OrderID uint64 `json:"order_id"`
}

type RefundMoney struct {
	OrderID uint64 `json:"order_id"`
}
