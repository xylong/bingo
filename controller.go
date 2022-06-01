package bingo

// Controller 控制器
type Controller interface {
	Route(group *Group)
}
