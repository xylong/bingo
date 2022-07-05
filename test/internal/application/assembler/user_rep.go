package assembler

import (
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/domain/model/user"
)

// UserRep 用户响应
type UserRep struct {
}

// M2D_SimpleUser 模型转dto
func (r *UserRep) M2D_SimpleUser(user *user.User) *dto.SimpleUser {
	return &dto.SimpleUser{
		ID:       user.ID,
		Avatar:   user.Avatar,
		Nickname: user.Nickname,
		Email:    user.Email,
	}
}
