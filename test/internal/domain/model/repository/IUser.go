package repository

type IUserRepo interface {
	GetByID(IModel) error
}
