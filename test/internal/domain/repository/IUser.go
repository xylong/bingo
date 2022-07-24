package repository

// IUser 用户接口
type IUser interface {
	Create(Modeler) error
	Get(interface{}) (int64, error)
}
