package model

import "fmt"

type Customer struct {
	CustomerID   string
	CustomerName string
	PhoneNumber  string
}

func (c Customer) String() string {
	return fmt.Sprintf("ID: %v, Name: %v, Phone Number: %v", c.CustomerID, c.CustomerName, c.PhoneNumber)
}
