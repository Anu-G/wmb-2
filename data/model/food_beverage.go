package model

import "fmt"

type FoodBeverage struct {
	FoodBeverageID string
	MenuName       string
	Price          float64
}

func (f FoodBeverage) String() string {
	return fmt.Sprintf("ID: %v, Name: %v, Price: %.2f", f.FoodBeverageID, f.MenuName, f.Price)
}
