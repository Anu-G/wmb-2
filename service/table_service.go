package service

import (
	"wmb_2/data/model"
	"wmb_2/data/repository"
)

type ServiceTableInterface interface {
	ViewTable() []model.Table
	GetTable(tableNo string) model.Table
	UpdateTable(id string)
}

type serviceTable struct {
	tableRepo repository.TableRepository
}

func (c *serviceTable) ViewTable() []model.Table {
	tables := c.tableRepo.FindByAvailability()
	// fmt.Println(tables)
	return tables
}

func (c *serviceTable) GetTable(tableNo string) model.Table {
	tables := c.tableRepo.FindById(tableNo)
	// fmt.Println(tables)
	return tables
}

func (c *serviceTable) UpdateTable(id string) {
	c.tableRepo.UpdateAvailability(id)
}

func NewServiceTable(tableRepo repository.TableRepository) ServiceTableInterface {
	return &serviceTable{
		tableRepo: tableRepo,
	}
}
