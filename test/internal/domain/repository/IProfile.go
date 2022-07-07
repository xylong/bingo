package repository

// IProfile 用户信息接口
type IProfile interface {
	GetByUser(model IModel) error
	Create(IModel) error
}
