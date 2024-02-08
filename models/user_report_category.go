package models

import (
	"greebel.core.be/core"
)

type (
	UserReportCategory struct {
		core.Model
		Category    string `json:"category" gorm:"column:category"`
		Description string `json:"description" gorm:"column:description"`
	}
)

func (UserReportCategory) TableName() string {
	return "user_report_categories"
}

func (p *UserReportCategory) Create() error {
	err := core.Create(&p)
	return err
}

func (p *UserReportCategory) Save() error {
	err := core.Save(&p)
	return err
}

func (p *UserReportCategory) Delete() error {
	err := core.Delete(&p)
	return err
}

func (p *UserReportCategory) FindbyID(id int) error {
	err := core.FindbyID(&p, id)
	return err
}

func (p *UserReportCategory) Find(filter interface{}) error {
	err := core.Find(&p, filter)
	return err
}

func (b *UserReportCategory) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result core.PagedFindResult, err error) {
	Question := []UserReportCategory{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = core.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}
