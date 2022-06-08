package repository

import (
	"fmt"

	"wmb_2/data/model"
)

type TableRepository interface {
	UpdateAvailability(ids ...string)
	FindByAvailability() []model.Table
	FindById(id string) model.Table
}

type tableRepository struct {
	db []model.Table
}

func (t *tableRepository) UpdateAvailability(ids ...string) {
	for _, id := range ids {
		for i, tbl := range t.db {
			if tbl.TableID == id {
				// tbl.IsAvailable = !tbl.IsAvailable
				t.db[i].IsAvailable = !tbl.IsAvailable
				break
			}
		}
	}
}

func (t *tableRepository) FindByAvailability() []model.Table {
	var tableAvailable []model.Table
	for _, tbl := range t.db {
		if tbl.IsAvailable {
			tableAvailable = append(tableAvailable, tbl)
		}
	}
	return tableAvailable
}

func (t *tableRepository) FindById(id string) model.Table {
	var selectedTable model.Table
	for i, tbl := range t.db {
		if tbl.TableID == id {
			selectedTable = t.db[i]
			break
		}
	}
	return selectedTable
}

func NewTableRepository(tableCapacity int) TableRepository {
	tableDb := make([]model.Table, tableCapacity)

	for i, data := range tableDb {
		data.TableID = fmt.Sprintf("T%02d", i+1)
		data.IsAvailable = true
		tableDb[i] = data
	}

	repo := new(tableRepository)
	repo.db = tableDb
	return repo
}
