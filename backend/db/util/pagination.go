package util

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// Paginate provides a solution to do both pagination and total number of
// records from one SELECT.
// See: http://andreyzavadskiy.com/2016/12/03/pagination-and-total-number-of-rows-from-one-select/
// Paginate also accepts extra select clauses, and default it selects all if
// `selects` is empty.
func Paginate(pg *Pagination, selects ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		clauses := []string{}
		if len(selects) == 0 {
			clauses = append(clauses, "*")
		} else {
			clauses = append(clauses, selects...)
		}
		clauses = append(clauses, "count(*) over () as total")
		return db.Select(clauses).Offset(pg.Offset).Limit(pg.Limit)
	}
}
