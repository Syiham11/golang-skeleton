package models

import (
	"greebel.core.be/core"
)

type (
	UserBanned struct {
		core.Model
		UserID    int    `json:"user_id" gorm:"column:user_id"`
		Status    int    `json:"status" gorm:"column:status"`
		Reason    string `json:"reason" gorm:"column:reason"`
		StartDate string `json:"start_date" gorm:"column:start_date"`
		EndDate   string `json:"end_date" gorm:"column:end_date"`
	}
)

func (UserBanned) TableName() string {
	return "users_banned"
}
