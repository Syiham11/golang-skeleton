package models

import (
	"greebel.core.be/core"
)

type (
	NewsCategory struct {
		core.Model
		Category string `json:"category" gorm:"column:category"`
	}
)

func (NewsCategory) TableName() string {
	return "news_category"
}

func (b *NewsCategory) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result core.PagedFindResult, err error) {
	Question := []NewsCategory{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = core.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}
