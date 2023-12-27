package core

type (
	ModelProduct struct {
		IdProduct int `json:"id_product" sql:"AUTO_INCREMENT" gorm:"primary_key,column:id_product"`
		// CreatedAt time.Time `json:"created_at" gorm:"column:created_at;DEFAULT:CURRENT_TIMESTAMP" sql:"DEFAULT:CURRENT_TIMESTAMP"`
		// UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" sql:"sql:"DEFAULT:CURRENT_TIMESTAMP"`
	}
)
