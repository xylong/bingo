package assembler

import (
	"github.com/go-playground/validator/v10"
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"github.com/xylong/bingo/test/internal/domain/model/userLog"
)

// UserReq 用户请求
type UserReq struct {
	v *validator.Validate
}

func (r *UserReq) D2M_User(req *dto.SimpleUserReq) *user.User {
	return user.New(user.WithID(req.ID))
}

func (r *UserReq) D2M_Log(u *user.User, logType int, remark string) *userLog.UserLog {
	return userLog.New(
		userLog.WithUserID(u.ID),
		userLog.WithType(uint8(logType)),
		userLog.WithRemark(remark))
}
