package models

import (
	"greebel.core.be/core"
)

type (
	OTPNoAuth struct {
		core.Model
		Code     string `json:"code" gorm:"column:code"`
		Category string `json:"category" gorm:"column:category"`
		Expired  int    `json:"expired" gorm:"column:expired"`
		Platform string `json:"platform" gorm:"column:platform"`
		Identity string `json:"identity" gorm:"column:identity"`
	}
)

func (OTPNoAuth) TableName() string {
	return "otp_no_auth"
}

func (p *OTPNoAuth) Create() error {
	err := core.Create(&p)
	return err
}

func (p *OTPNoAuth) Save() error {
	err := core.Save(&p)
	return err
}

func (p *OTPNoAuth) Delete() error {
	err := core.Delete(&p)
	return err
}

func (p *OTPNoAuth) FindbyID(id int) error {
	err := core.FindbyID(&p, id)
	return err
}

func (p *OTPNoAuth) Find(filter interface{}) error {
	err := core.Find(&p, filter)
	return err
}

func (b *OTPNoAuth) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result core.PagedFindResult, err error) {
	Question := []OTPNoAuth{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = core.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}
