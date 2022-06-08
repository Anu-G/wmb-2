package service

import (
	"fmt"

	"wmb_2/data/model"
	"wmb_2/data/repository"
)

type ServiceMenuInterface interface {
	FindMenuById(id string) model.FoodBeverage
	FindMenuByName(keyword string) []model.FoodBeverage
}

type serviceMenu struct {
	fnbRepo repository.FnBRepository
}

func (c *serviceMenu) FindMenuById(id string) model.FoodBeverage {
	fnb := c.fnbRepo.FindById(id)
	if fnb.FoodBeverageID == "" {
		fmt.Printf("Product with ID %s not found\n", id)
	} else {
		fmt.Printf("Product found : [%v]\n", fnb)
	}
	return fnb
}

func (c *serviceMenu) FindMenuByName(keyword string) []model.FoodBeverage {
	fnbs := c.fnbRepo.FindByName(keyword)
	if len(fnbs) == 0 {
		fmt.Printf("Product with keyword %s not found\n", keyword)
	} else {
		fmt.Printf("Product found: [%v]\n", fnbs)
	}
	return fnbs
}

func NewServiceMenu(fnbRepo repository.FnBRepository) ServiceMenuInterface {
	return &serviceMenu{
		fnbRepo: fnbRepo,
	}
}
