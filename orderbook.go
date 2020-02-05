package orderbook

const maxPrice = 10000000

// OrderStatus represents the status of an individual order.
type OrderStatus int

// The valid statuses of an order on the book.
const (
	OrderStatusNew OrderStatus = iota
	OrderStatusOpen
	OrderStatusPartial
	OrderStatusFilled
	OrderStatusCanceled
)

// Order represents an order on the book. Orders are either buy or sell.
// Each order is linted to the next order at the same price point.
type Order struct {
	id     uint64
	isBuy  bool
	price  uint32
	amount uint32
	status OrderStatus
	next   *Order
}

// OrderBook keeps track of the current maximum bid and minimum ask,
// the orders, and possible price points.
type OrderBook struct {
	ask        uint32
	bid        uint32
	orderIndex map[uint64]*Order
	prices     [maxPrice]*PricePoint
	actions    chan<- *Action
}

// OpenOrder inserts a new order into the book.
// It appends the order to the list of orders at its price point,
// updates the bid or ask, and adds an entry in the order index.
func (orderBook *OrderBook) OpenOrder(order *Order) {
	pricePoint := orderBook.prices[order.price]
	pricePoint.Insert(order)

	order.status = OrderStatusNew

	if order.isBuy && order.price > orderBook.bid {
		orderBook.bid = order.price
	}

	if !order.isBuy && order.price < orderBook.ask {
		orderBook.ask = order.price
	}

	orderBook.orderIndex[order.id] = order
}

// CancelOrder cancels an order by setting its amount to 0 and
// status to OrderStatusCanceled.
func (orderBook *OrderBook) CancelOrder(id uint64) {
	if order, ok := orderBook.orderIndex[id]; ok {
		order.amount = 0
		order.status = OrderStatusCanceled
	}

	orderBook.actions <- NewCanceledAction(id)
}
