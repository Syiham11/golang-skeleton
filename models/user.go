package models

import (
	"greebel.core.be/core"
)

type (
	User struct {
		core.Model
		Name           string            `json:"name" gorm:"column:name"`
		Username       string            `json:"username" gorm:"column:username"`
		Email          string            `json:"email" gorm:"column:email"`
		Password       string            `json:"password" gorm:"column:password"`
		Address        string            `json:"address" gorm:"column:address"`
		Company        string            `json:"company" gorm:"column:company"`
		Paket          int               `json:"paket" gorm:"column:paket"`
		PhoneNumber    string            `json:"phone_number" gorm:"column:phone_number"`
		StatusActive   int               `json:"status_active" gorm:"column:status_active"`
		IsPartner      int               `json:"is_partner" gorm:"column:is_partner"`
		ProfilePicture string            `json:"profile_picture" gorm:"column:profile_picture"`
		BankAccounts   []UserBankAccount `json:"bank_accounts" gorm:"foreignkey:user_id"`
		PlayerID       string            `json:"player_id" gorm:"column:player_id"`
	}
)

func (User) TableName() string {
	return "users"
}

func (p *User) Create() error {
	err := core.Create(&p)
	return err
}

func (p *User) Save() error {
	err := core.Save(&p)
	return err
}

func (p *User) Delete() error {
	err := core.Delete(&p)
	return err
}

func (p *User) FindbyID(id int) error {
	err := core.FindbyID(&p, id)
	return err
}

func (p *User) Find(filter interface{}) error {
	err := core.Find(&p, filter)
	return err
}

func (b *User) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result core.PagedFindResult, err error) {
	Question := []User{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = core.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}
