package repository

// IProfileRepo 用户信息接口
type IProfileRepo interface {
	GetByUser(model IModel) error
}
