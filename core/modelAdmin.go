package core

import (
	"time"
)

type (
	ModeAdmin struct {
		IdAdm     int       `json:"id_adm" sql:"AUTO_INCREMENT" gorm:"primary_key,column:id_adm"`
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at;DEFAULT:CURRENT_TIMESTAMP" sql:"DEFAULT:CURRENT_TIMESTAMP"`
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" sql:"sql:"DEFAULT:CURRENT_TIMESTAMP"`
	}
)
