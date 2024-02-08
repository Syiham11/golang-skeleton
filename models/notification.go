package models

import (
	"encoding/json"
	"fmt"
	"greebel.core.be/core"
	"greebel.core.be/helper"
)

type (
	Notification struct {
		core.Model
		UserID            int             `json:"user_id" gorm:"column:user_id"`
		User              *User           `json:"user" gorm:"foreignkey:user_id"`
		NotificationType  string          `json:"notification_type" gorm:"column:notification_type"`
		Content           string          `json:"content" gorm:"column:content"`
		URL               string          `json:"url" gorm:"column:url"`
		AdditionalData    string          `json:"_" gorm:"column:additional_data"`
		AdditionalDataArr json.RawMessage `json:"additional_data" gorm:"-"`
		FromUserID        int             `json:"from_user_id" gorm:"column:from_user_id"`
		FromUser          *User           `json:"from_user" gorm:"foreignkey:from_user_id"`
		ReadStatus        int             `json:"read_status" gorm:"column:read_status"`

		Heading           string                 `json:"heading" gorm:"-"`
		Segments          []string               `json:"segments" gorm:"-"`
		ExternalUserIds   []string               `json:"external_user_ids" gorm:"-"`
		AdditionalDataMap map[string]interface{} `json:"additional_data_map" gorm:"-"`
	}
)

func (Notification) TableName() string {
	return "notifications"
}

func (p *Notification) AfterCreate() error {
	err := helper.SendNotification(p.Heading, p.Content, p.URL, p.Segments, p.ExternalUserIds, p.AdditionalDataMap)
	if err != nil {
		return err
	}
	return nil
}

func (p *Notification) AfterFind() error {

	in := []byte(p.AdditionalData)

	err := json.Unmarshal(in, &p.AdditionalDataArr)
	if err != nil {
		panic(err)
	}

	// Context is []byte, so you can keep it as string in DB
	//	fmt.Println("ctx:", string(p.AdditionalDataArr))

	// Marshal back to json (as original)
	out, _ := json.Marshal(&p.AdditionalDataArr)
	//	fmt.Println(string(out))
	p.AdditionalDataArr = out
	return nil
}

func (Notification) GetNotifications(page int, rows int, orderby string, sort string, filter interface{}) ([]Notification, int, error) {
	query := core.App.DB.Table("notifications").
		Preload("User").
		Preload("FromUser")
	query = core.ConditionQuery(query, filter)
	query = query.Order(fmt.Sprintf("%s %s", orderby, sort))

	temp := query
	notifications := []Notification{}
	var (
		totalRows int
		offset    int
	)

	temp.Find(&notifications).Count(&totalRows)

	if rows > 0 {
		offset = (page * rows) - rows
		query = query.Limit(rows).Offset(offset)
	}

	err := query.Find(&notifications).Error

	return notifications, totalRows, err
}

func (b *Notification) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result core.PagedFindResult, err error) {
	Question := []Notification{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = core.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}

func (p *Notification) FindbyID(id int) error {
	err := core.FindbyID(&p, id)
	return err
}

func (p *Notification) Find(filter interface{}) error {
	err := core.Find(&p, filter)
	return err
}

func (p *Notification) Save() error {
	err := core.Save(&p)
	return err
}
