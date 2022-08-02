package assembler

import (
	"github.com/xylong/bingo/test/internal/application/dto"
	"github.com/xylong/bingo/test/internal/domain/aggregation"
	"github.com/xylong/bingo/test/internal/domain/model/user"
	"github.com/xylong/bingo/test/internal/domain/model/userLog"
	"github.com/xylong/bingo/test/internal/infrastructure/dao/GormDao"
	"github.com/xylong/bingo/test/internal/lib/db"
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
		Phone:    user.HidePhone(),
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

func (r *UserRep) M2D_UserInfo(req *dto.UserLogReq, member *aggregation.Member) *dto.UserInfo {
	member.User.ID = req.ID
	if err := member.GetUser(); err != nil {
		return nil
	}

	info := &dto.UserInfo{
		ID:       member.User.ID,
		Nickname: member.User.Nickname,
	}

	member.Log = userLog.New(userLog.WithUserID(member.User.ID))
	member.Log.Dao = GormDao.NewUserLogDao(db.DB)

	info.Log = r.M2D_UserLogs(member.GetLog(req))
	return info
}

func (r *UserRep) M2D_UserLogs(logs []*userLog.UserLog) (userLogs []*dto.UserLog) {
	if len(logs) > 0 {
		for _, log := range logs {
			userLogs = append(userLogs, &dto.UserLog{
				ID:   log.ID,
				Log:  log.Remark,
				Date: log.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}
	}

	return userLogs
}
