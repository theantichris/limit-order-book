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

// NewCanceledAction sends a canceled action to the order book.
func NewCanceledAction(id uint64) *Action {
	return &Action{
		ActionType: ActionTypeCanceled,
		OrderID:    id,
	}
}
