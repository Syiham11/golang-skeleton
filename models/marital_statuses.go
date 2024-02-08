package models

import (
	"greebel.core.be/core"
)

type (
	MaritalStatus struct {
		core.Model
		Name string `json:"name" gorm:"column:name"`
	}
)

func (MaritalStatus) TableName() string {
	return "marital_statuses"
}

func (b *MaritalStatus) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result core.PagedFindResult, err error) {
	Question := []MaritalStatus{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = core.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}
