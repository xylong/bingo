package repository

// Profiler 用户信息接口
type Profiler interface {
	Create(Modeler) error
}
