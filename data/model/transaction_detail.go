package model

import "fmt"

type TransactionDetail struct {
	TransactionDetailID string
	Order               CustomerOrder
}

func (t TransactionDetail) String() string {
	return fmt.Sprintf("ID: %v, Quantity: %d, Menu: %v", t.TransactionDetailID, t.Order.Quantity, t.Order.OrderedMenu)
}
