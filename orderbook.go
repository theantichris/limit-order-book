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
	isBuy  bool        // true for buy, false for sell
	price  uint32      // the price per unit to buy or sell at
	amount uint32      // the amount the user wants to buy for the price
	status OrderStatus // status of the order
	next   *Order      // the next order on the book
}

// OrderBook keeps track of the order book.
type OrderBook struct {
	ask        uint32                // minimum ask
	bid        uint32                // maximum bid
	orderIndex map[uint64]*Order     // list of orders on the book
	prices     [maxPrice]*PricePoint // list of price points
	actions    chan<- *Action        // channel for order book operations
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

// FillSell fills a sell order.
// Starts at the maximum bid and iterates over all open orders until
// it fills the order or exhausts the open orders at the price point.
// It then moves down to the next price point and repeats.
func (orderBook *OrderBook) FillSell(order *Order) {
	for orderBook.bid >= order.price && order.amount > 0 {
		pricePoint := orderBook.prices[orderBook.bid]
		orderHead := pricePoint.orderHead

		for orderHead != nil {
			orderBook.fillOrder(order, orderHead)
			orderHead = orderHead.next
			pricePoint.orderHead = orderHead
		}

		orderBook.bid--
	}
}

func (orderBook *OrderBook) fillOrder(order, orderHead *Order) {
	if orderHead.amount >= order.amount {
		orderBook.actions <- NewFilledAction(order, orderHead)
		orderHead.amount -= order.amount

		order.amount = 0
		order.status = OrderStatusFilled

		return
	}

	if orderHead.amount > 0 {
		orderBook.actions <- NewPartialFilledAction(order, orderHead)

		order.amount -= orderHead.amount
		order.status = OrderStatusPartial

		orderHead.amount = 0
	}
}
