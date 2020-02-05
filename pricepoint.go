package orderbook

// PricePoint represents a discreet limit price.
// Contains pointers to the first and last order entered at that price.
type PricePoint struct {
	orderHead *Order
	orderTail *Order
}

// Insert adds an order to the current price point.
func (pricePoint *PricePoint) Insert(order *Order) {
	if pricePoint.orderHead == nil {
		pricePoint.orderHead = order
		pricePoint.orderTail = order
	} else {
		pricePoint.orderTail.nextOrder = order
		pricePoint.orderTail = order
	}
}
