package GormDao

import "github.com/xylong/bingo/test/internal/domain/model/userLog"

type UserLogDao struct {
}

func (d *UserLogDao) GetByUser(uid int) []*userLog.UserLog {
	return nil
}

func (d *UserLogDao) Save(log *userLog.UserLog) error {
	return nil
}
