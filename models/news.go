package models

import (
	"greebel.core.be/core"
)

type (
	News struct {
		core.Model
		Title          string        `json:"title" gorm:"column:title"`
		NewsCategoryID int           `json:"news_category_id" gorm:"column:news_category_id"`
		Content        string        `json:"content" gorm:"column:content"`
		Tags           string        `json:"tags" gorm:"column:tags"`
		Thumbnail      string        `json:"thumbnail" gorm:"column:thumbnail"`
		AuthorName     string        `json:"author_name" gorm:"column:author_name"`
		AuthorImage    string        `json:"author_image" gorm:"column:author_image"`
		TotalViews     int           `json:"total_views" gorm:"column:total_views"`
		Source         string        `json:"source" gorm:"column:source"`
		SlugURL        string        `json:"slug_url" gorm:"column:slug_url"`
		NewsCategory   *NewsCategory `json:"news_category" gorm:"foreignkey:news_category_id"`
	}
)

func (News) GetNews(page, rows int, filter interface{}) ([]News, int, error) {
	query := core.App.DB.Table("news").
		Preload("NewsCategory")
	query = core.ConditionQuery(query, filter)

	temp := query
	news := []News{}
	var (
		totalRows int
		offset    int
	)

	temp.Find(&news).Count(&totalRows)

	if rows > 0 {
		offset = (page * rows) - rows
		query = query.Limit(rows).Offset(offset)
	}

	err := query.Find(&news).Error

	return news, totalRows, err
}

func (News) TableName() string {
	return "news"
}

func (p *News) FindbyID(id int) error {
	err := core.FindbyID(&p, id)
	return err
}

func (b *News) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result core.PagedFindResult, err error) {
	Question := []News{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = core.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}
