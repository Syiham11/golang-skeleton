package models

import (
	"greebel.core.be/core"
)

type (
	City struct {
		core.Model
		Name string `json:"name" gorm:"column:name"`
	}
)

func (City) TableName() string {
	return "cities"
}

func (b *City) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result core.PagedFindResult, err error) {
	Question := []City{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = core.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}

func (b *City) PageList(page, rows int, partnerType string, filter interface{}) ([]City, int, error) {
	query := core.App.DB.Table("partners").
		Select(`
		cities.name AS name,
		cities.id AS id,
		count(cities.name) as total,
		(
				CASE
						WHEN talents.id IS NOT NULL THEN "TALENT"
						WHEN influencers.id IS NOT NULL THEN "INFLUENCER"
						WHEN venues.id IS NOT NULL THEN "VENUE"
						WHEN vendors.id IS NOT NULL THEN "VENDOR"
						ELSE "UNKNOWN"
				END
		) AS partner_type,
		(
				CASE
						WHEN talents.id IS NOT NULL THEN talents.price_rate
						ELSE NULL
				END
		) AS price,
		(
				CASE
						WHEN talents.id IS NOT NULL THEN talents.price_discount
						ELSE NULL
				END
		) AS price_discount,
		(
				CASE
						WHEN talents.id IS NOT NULL THEN ((talents.price_discount / talents.price_rate) * 100)
						ELSE NULL
				END
		) AS discount
		`).
		Joins("JOIN users ON users.id = partners.user_id").
		Joins("LEFT JOIN cities ON cities.id = partners.city_id").
		Joins("LEFT JOIN partner_categories pc ON partners.category_id = pc.id").
		Joins("LEFT JOIN talents ON talents.partner_id = partners.id").
		Joins("LEFT JOIN influencers ON influencers.partner_id = partners.id").
		Joins("LEFT JOIN venues ON venues.partner_id = partners.id").
		Joins("LEFT JOIN vendors ON vendors.partner_id = partners.id").
		Where("partners.approval_status = ?", "1").
		Where("cities.name != ?", "").
		Group("cities.name").
		Limit("10").
		Order("total DESC")

	query = core.ConditionQuery(query, filter)

	switch partnerType {
	case "VENUE":
		query = query.Joins("JOIN venues b ON partners.id = b.partner_id")

	case "VENDOR":
		query = query.Joins("JOIN vendors b ON partners.id = b.partner_id")

	case "INFLUENCER":
		query = query.Joins("JOIN influencers b ON partners.id = b.partner_id")

	case "TALENT":
		query = query.Joins("JOIN talents b ON partners.id = b.partner_id")
	}

	temp := query
	partners := []City{}
	var (
		totalRows int
		offset    int
	)

	temp.Find(&partners).Count(&totalRows)

	if rows > 0 {
		offset = (page * rows) - rows
		query = query.Limit(rows).Offset(offset)
	}

	err := query.Find(&partners).Error

	return partners, totalRows, err
}

func (p *City) FindbyID(id int) error {
	err := core.FindbyID(&p, id)
	return err
}
