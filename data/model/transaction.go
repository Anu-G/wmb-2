package model

import (
	"fmt"
	"time"
)

type Transaction struct {
	TransactionID     string
	TransactionDate   time.Time
	FinishTransaction bool

	CustomerID        Customer
	TableID           Table
	TransactionDetail []TransactionDetail
}

func (t Transaction) String() string {
	return fmt.Sprintf(`
	ID: %v
	Table: [%v]
	Transaction Date: %v
	Customer: [%v]
	Transaction Detail: [%v]
	Paid? %v
	`,
		t.TransactionID, t.TableID, t.TransactionDate, t.CustomerID, t.TransactionDate, t.FinishTransaction)
}
