package model

import "fmt"

type Table struct {
	TableID     string
	IsAvailable bool
}

func (t Table) String() string {
	return fmt.Sprintf("Table Number: %v, Availability: %v", t.TableID, t.IsAvailable)
}
