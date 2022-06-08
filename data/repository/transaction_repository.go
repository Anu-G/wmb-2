package repository

import (
	"time"

	"wmb_2/data/model"
	"wmb_2/utils"
)

type TrxRepository interface {
	Create(customer model.Customer, tableNo model.Table, orders []model.CustomerOrder) string
	UpdateBySettled(billNo string)
	FindById(billNo string) model.Transaction
}

type trxRepository struct {
	db []model.Transaction
}

func (t *trxRepository) FindById(billNo string) model.Transaction {
	var billSelected model.Transaction
	for _, bill := range t.db {
		if bill.TransactionID == billNo {
			billSelected = bill
			break
		}
	}
	return billSelected
}

func (t *trxRepository) Create(customer model.Customer, table model.Table, orders []model.CustomerOrder) string {
	var billDetails []model.TransactionDetail

	for _, order := range orders {
		billDetails = append(billDetails, model.TransactionDetail{
			TransactionDetailID: utils.GenerateId(),
			Order:               order,
		})
	}
	newBillNo := utils.GenerateId()
	newBill := model.Transaction{
		TransactionID:     newBillNo,
		TableID:           table,
		TransactionDate:   time.Now(),
		CustomerID:        customer,
		TransactionDetail: billDetails,
	}
	t.db = append(t.db, newBill)
	return newBillNo
}

func (t *trxRepository) UpdateBySettled(billNo string) {
	for idx, bill := range t.db {
		if bill.TransactionID == billNo {
			t.db[idx].FinishTransaction = true
			break
		}
	}
}

func NewTrxRepository() TrxRepository {
	repo := new(trxRepository)
	return repo
}
