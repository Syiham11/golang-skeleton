package models

import (
	"greebel.core.be/core"
)

type (
	OTP struct {
		core.Model
		UserID   int    `json:"user_id" gorm:"column:user_id"`
		Code     string `json:"code" gorm:"column:code"`
		Category string `json:"category" gorm:"column:category"`
		Expired  int    `json:"expired" gorm:"column:expired"`
		TimerOtp int    `json:"timer_otp" gorm:"column:timer_otp"`
		Platform string `json:"platform" gorm:"column:platform"`
		Identity string `json:"indentity" gorm:"column:indentity"`
		Used     string `json:"used" gorm:"column:used"`
	}
)

func (OTP) TableName() string {
	return "otp"
}

func (p *OTP) Create() error {
	err := core.Create(&p)
	return err
}

func (p *OTP) Save() error {
	err := core.Save(&p)
	return err
}

func (p *OTP) Delete() error {
	err := core.Delete(&p)
	return err
}

func (p *OTP) FindbyID(id int) error {
	err := core.FindbyID(&p, id)
	return err
}

func (p *OTP) Find(filter interface{}) error {
	err := core.Find(&p, filter)
	return err
}

func (b *OTP) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result core.PagedFindResult, err error) {
	Question := []OTP{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = core.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}
