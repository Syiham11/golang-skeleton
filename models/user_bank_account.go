package models

import (
	"greebel.core.be/core"
)

type (
	UserBankAccount struct {
		core.Model
		BankName      string `json:"bank_name" gorm:"column:bank_name"`
		UserID        int    `json:"user_id" gorm:"column:user_id"`
		BankID        int    `json:"bank_id" gorm:"column:bank_id"`
		AccountNumber string `json:"account_number" gorm:"column:account_number"`
		AccountName   string `json:"account_name" gorm:"column:account_name"`
	}
)

func (UserBankAccount) TableName() string {
	return "user_bank_accounts"
}

func (b *UserBankAccount) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result core.PagedFindResult, err error) {
	Question := []UserBankAccount{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = core.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}

func (p *UserBankAccount) FindbyID(id int) error {
	err := core.FindbyID(&p, id)
	return err
}

func (p *UserBankAccount) Delete() error {
	err := core.Delete(&p)
	return err
}
