package assembler

import (
	"github.com/go-playground/validator/v10"
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/domain/model/user"
)

// UserReq 用户请求
type UserReq struct {
	v *validator.Validate
}

func (r *UserReq) D2M_User(req *dto.SimpleUserReq) *user.User {
	if err := r.v.Struct(req); err != nil {
		panic(err.Error())
	}

	return user.New(user.WithID(req.ID))
}
