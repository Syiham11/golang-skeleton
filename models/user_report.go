package models

import (
	"greebel.core.be/core"
)

type (
	UserReport struct {
		core.Model
		ReporterID       int `json:"reporter_id" gorm:"column:reporter_id"`
		UserID           int `json:"user_id" gorm:"column:user_id"`
		ReportCategoryID int `json:"report_category_id" gorm:"column:report_category_id"`
	}
)

func (UserReport) TableName() string {
	return "user_reports"
}

func (p *UserReport) Create() error {
	err := core.Create(&p)
	return err
}

func (p *UserReport) Save() error {
	err := core.Save(&p)
	return err
}

func (p *UserReport) Delete() error {
	err := core.Delete(&p)
	return err
}

func (p *UserReport) FindbyID(id int) error {
	err := core.FindbyID(&p, id)
	return err
}

func (p *UserReport) Find(filter interface{}) error {
	err := core.Find(&p, filter)
	return err
}

func (b *UserReport) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result core.PagedFindResult, err error) {
	Question := []UserReport{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = core.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}
