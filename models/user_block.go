package models

import (
	"greebel.core.be/core"
)

type (
	UserBlock struct {
		core.Model
		UserID        int `json:"user_id" gorm:"column:user_id"`
		BlockedUserID int `json:"blocked_user_id" gorm:"column:blocked_user_id"`
	}
)

func (UserBlock) TableName() string {
	return "users_block"
}
