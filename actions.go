package orderbook

// ActionType indicates the type of an action.
type ActionType string

// The valid types of actions.
const (
	ActionTypeBuy           = "BUY"
	ActionTypeSell          = "SELL"
	ActionTypeCancel        = "CANCEL"
	ActionTypeCanceled      = "CANCELED"
	ActionTypePartialFilled = "PARTIAL_FILLED"
	ActionTypeFilled        = "FILLED"
	ActionTypeDone          = "DONE"
)

// Action represents an operation on the order book.
type Action struct {
	ActionType  ActionType `json:"actionType"`
	OrderID     uint64     `json:"orderId"`
	FromOrderID uint64     `json:"fromOrderId"`
	Amount      uint32     `json:"amount"`
	Price       uint32     `json:"price"`
}

// NewCanceledAction creates a canceled action.
func NewCanceledAction(id uint64) *Action {
	return &Action{
		ActionType: ActionTypeCanceled,
		OrderID:    id,
	}
}

// NewFilledAction creates a filled action.
func NewFilledAction(order *Order, fromOrder *Order) *Action {
	return &Action{
		ActionType:  ActionTypeFilled,
		OrderID:     order.id,
		FromOrderID: fromOrder.id,
		Amount:      order.amount,
		Price:       fromOrder.price,
	}
}

// NewPartialFilledAction creates a partial filled action.
func NewPartialFilledAction(order *Order, fromOrder *Order) *Action {
	return &Action{
		ActionType:  ActionTypePartialFilled,
		OrderID:     order.id,
		FromOrderID: fromOrder.id,
		Amount:      fromOrder.amount,
		Price:       fromOrder.price,
	}
}
