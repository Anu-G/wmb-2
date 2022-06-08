package service

import (
	"fmt"

	"wmb_2/data/model"
	"wmb_2/data/repository"
)

type ServiceTransactionInterface interface {
	TakeOrder(customer model.Customer, tableNo string, orders []model.CustomerOrder) string
	OrderPayment(billNo string)
}

type serviceTransaction struct {
	trxRepo      repository.TrxRepository
	serviceTable ServiceTableInterface
}

func (c *serviceTransaction) TakeOrder(customer model.Customer, tableNo string, orders []model.CustomerOrder) string {
	var newBillNo string
	tableReserve := c.serviceTable.GetTable(tableNo)
	if tableReserve.IsAvailable {
		newBillNo = c.trxRepo.Create(customer, tableReserve, orders)
		// c.tableRepo.UpdateAvailability(tableReserve.TableID)
		c.serviceTable.UpdateTable(tableReserve.TableID)
		fmt.Printf("Order %s successfully created\n", newBillNo)
	} else {
		fmt.Printf("Table %s is not available\n", tableReserve.TableID)
	}
	return newBillNo
}

func (c *serviceTransaction) OrderPayment(billNo string) {
	bill := c.trxRepo.FindById(billNo)
	if bill.TransactionID != "" {
		bill.FinishTransaction = true
	} else {
		fmt.Println("No order have been made", billNo)
		return
	}
	// c.tableRepo.UpdateAvailability(bill.TableNo.TableNo)
	c.serviceTable.UpdateTable(bill.TableID.TableID)
	fmt.Printf("Order %s successfully paid\n", billNo)
}

func NewServiceTransaction(trxRepo repository.TrxRepository, serviceTable ServiceTableInterface) ServiceTransactionInterface {
	return &serviceTransaction{
		trxRepo:      trxRepo,
		serviceTable: serviceTable,
	}
}
