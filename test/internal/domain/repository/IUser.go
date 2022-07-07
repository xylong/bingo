package repository

import "github.com/xylong/bingo/test/internal/domain/model/user"

// IUser 用户接口
type IUser interface {
	Get(*user.User) error
	GetByPhone(string) *user.User
	Create(*user.User) error
}
