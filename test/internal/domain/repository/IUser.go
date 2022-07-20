package repository

// IUser 用户接口
type IUser interface {
	Create(Modeler) error
}
