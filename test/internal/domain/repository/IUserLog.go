package repository

import "github.com/xylong/bingo/test/internal/domain/model/userLog"

type IUserLog interface {
	GetByUser(int) []*userLog.UserLog
	Save(*userLog.UserLog) error
}
