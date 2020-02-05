package orderbook

const maxPrice = 10000000

// OrderStatus represents the status of an individual order.
type OrderStatus int

const (
	new OrderStatus = iota
	open
	partial
	filled
	cancelled
)

// Order represents an order on the book. Orders are either buy or sell.
// Each order is linted to the next order at the same price point.
type Order struct {
	id     uint64
	isBuy  bool
	price  uint32
	amount uint32
	next   *Order
}

// PricePoint represents a discreet limit price.
// Contains pointers to the first and last order entered at that price.
type PricePoint struct {
	orderHead *Order
	orderTail *Order
}

// OrderBook keeps track of the current maximum bid and minimum ask, the orders, and possible price points.
type OrderBook struct {
	ask        uint32
	bid        uint32
	orderIndex map[uint64]*Order
	prices     [maxPrice]*PricePoint
}
