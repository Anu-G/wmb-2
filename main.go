package main

import (
	"wmb_2/data/model"
	"wmb_2/data/repository"
	"wmb_2/service"
)

func main() {
	// // FNB
	// fnbrepo := repository.NewFnBRepository()

	// fnb06 := model.FoodBeverage{
	// 	FoodBeverageID: "F003",
	// 	MenuName:       "Nasi Rendang",
	// 	Price:          12000,
	// }
	// fnbrepo = fnbrepo.AddFnBRepository(&fnb06)
	// // fmt.Println(fnbrepo.FindById("f001"))
	// // fmt.Println(fnbrepo.FindByName("Jus"))

	// f1 := fnbrepo.FindById("F002")
	// b1 := fnbrepo.FindById("B002")
	// //

	// // Table
	// tablerepo := repository.NewTableRepository(20)
	// tablerepo.UpdateAvailability("T01", "T02", "T03")
	// tablerepo.UpdateAvailability("T01")
	// // fmt.Println(tablerepo.FindByAvailability())

	// table01 := tablerepo.FindById("T01")
	// table05 := tablerepo.FindById("T05")
	// //

	// // Customer
	// customer01 := model.Customer{
	// 	CustomerID:   "C00001",
	// 	PhoneNumber:  "08788123123",
	// 	CustomerName: "Jution",
	// }
	// //

	// // Transaction
	// trxRepo := repository.NewTrxRepository()
	// billNo1 := trxRepo.Create(customer01, table01, []model.CustomerOrder{
	// 	{OrderedMenu: f1, Quantity: 1},
	// 	{OrderedMenu: b1, Quantity: 2},
	// })
	// fmt.Println(trxRepo.FindById(billNo1))
	// billNo2 := trxRepo.Create(customer01, table05, []model.CustomerOrder{
	// 	{OrderedMenu: f1, Quantity: 1},
	// })
	// fmt.Println(trxRepo.FindById(billNo2))

	// trxRepo.UpdateBySettled(billNo1)
	// fmt.Println(trxRepo.FindById(billNo1))

	fnbRepo := repository.NewFnBRepository()
	tableRepo := repository.NewTableRepository(30)
	trxRepo := repository.NewTrxRepository()

	findMenuUseCase := service.NewServiceMenu(fnbRepo)
	tableViewUseCase := service.NewServiceTable(tableRepo)
	customerOrderUseCase := service.NewServiceTransaction(trxRepo, tableViewUseCase)

	f1 := findMenuUseCase.FindMenuById("F001")
	b1 := findMenuUseCase.FindMenuById("B001")
	findMenuUseCase.FindMenuByName("asi")

	customerPaymentUseCase := service.NewServiceTransaction(trxRepo, tableViewUseCase)
	customer01 := model.Customer{
		CustomerID:   "C00001",
		PhoneNumber:  "08788123123",
		CustomerName: "Jution",
	}

	newBillNo := customerOrderUseCase.TakeOrder(customer01, "T02", []model.CustomerOrder{
		{OrderedMenu: f1, Quantity: 1},
		{OrderedMenu: b1, Quantity: 2},
	})
	newBillNo2 := customerOrderUseCase.TakeOrder(customer01, "T02", []model.CustomerOrder{
		{OrderedMenu: f1, Quantity: 1},
		{OrderedMenu: b1, Quantity: 2},
	})
	customerPaymentUseCase.OrderPayment(newBillNo)
	customerPaymentUseCase.OrderPayment(newBillNo2)

	tableViewUseCase.ViewTable()
}
