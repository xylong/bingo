package dto

// Pagination 分页
type Pagination struct {
	Page     int `json:"page" form:"page" binding:"required,gt=0"`
	PageSize int `json:"page_size" form:"page_size" binding:"required,gt=0"`
}
