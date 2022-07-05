package repository

// IUserRepo 用户接口
type IUserRepo interface {
	GetByID(IModel) error
	Create(IModel) error
}
