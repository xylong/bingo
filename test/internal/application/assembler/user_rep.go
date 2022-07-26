package assembler

import (
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/domain/aggregation"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"github.com/xylong/bingo/test/internal/domain/model/userLog"
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
		Phone:    user.Phone,
	}
}

func (r *UserRep) M2D_SimpleList(users []*user.User) []*dto.SimpleUser {
	var list []*dto.SimpleUser

	if len(users) > 0 {
		for _, u := range users {
			list = append(list, &dto.SimpleUser{
				ID:       u.ID,
				Avatar:   u.Avatar,
				Nickname: u.Nickname,
				Phone:    u.HidePhone(),
			})
		}
	}

	return list
}

func (r *UserRep) M2D_UserInfo(member *aggregation.Member) *dto.UserInfo {
	info := &dto.UserInfo{
		ID:       member.User.ID,
		Nickname: member.User.Nickname,
	}
	info.Logs = r.M2D_UserLogs(member.GetLog())

	return info
}

func (r *UserRep) M2D_UserLogs(logs []*userLog.UserLog) []*dto.UserLog {
	return nil
}
