package repository

import (
	"strings"

	"wmb_2/data/model"
)

type FnBRepository interface {
	FindById(id string) model.FoodBeverage
	FindByName(name string) []model.FoodBeverage
	AddFnBRepository(fnb *model.FoodBeverage) FnBRepository
}

type fnbRepository struct {
	db []model.FoodBeverage
}

func (f *fnbRepository) FindById(id string) model.FoodBeverage {
	var fnbSelected model.FoodBeverage
	for _, fnb := range f.db {
		if strings.EqualFold(fnb.FoodBeverageID, id) {
			fnbSelected = fnb
			break
		}
	}
	return fnbSelected
}

func (f *fnbRepository) FindByName(name string) []model.FoodBeverage {
	var fnbSelected []model.FoodBeverage
	for _, fnb := range f.db {
		if strings.Contains(strings.ToLower(fnb.MenuName), strings.ToLower(name)) {
			fnbSelected = append(fnbSelected, fnb)
		}
	}
	return fnbSelected
}

func (f *fnbRepository) AddFnBRepository(fnb *model.FoodBeverage) FnBRepository {
	f.db = append(f.db, *fnb)
	return f
}

func NewFnBRepository() FnBRepository {
	repo := new(fnbRepository)
	fnb01 := model.FoodBeverage{
		FoodBeverageID: "F001",
		MenuName:       "Nasi Goreng",
		Price:          15000,
	}
	fnb02 := model.FoodBeverage{
		FoodBeverageID: "B001",
		MenuName:       "Es Teh Manis",
		Price:          4000,
	}
	fnb03 := model.FoodBeverage{
		FoodBeverageID: "F002",
		MenuName:       "Nasi Uduk",
		Price:          6000,
	}
	fnb04 := model.FoodBeverage{
		FoodBeverageID: "B002",
		MenuName:       "Es Kopi",
		Price:          4000,
	}
	repo.db = []model.FoodBeverage{fnb01, fnb02, fnb03, fnb04}
	return repo
}
