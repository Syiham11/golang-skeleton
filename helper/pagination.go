package helper

type Pagination struct {
	Page     int `query:"page"`
	PageSize int `query:"pageSize"`
}
