package static

const (
	PendingOrder = iota
	CancelledOrder
	PaidOrder
	Refunded
)

var OrderStatus = []string{"待支付订单", "取消订单", "已支付订单", "已退款"}
