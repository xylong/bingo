package repository

import (
	"github.com/xylong/bingo/test/internal/domain/model/profile"
	"github.com/xylong/bingo/test/internal/domain/model/user"
)

// IProfile 用户信息接口
type IProfile interface {
	GetByUser(*user.User) (*profile.Profile, error)
	Create(*profile.Profile) error
}
