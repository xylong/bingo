package aggregation

import (
	"github.com/xylong/bingo/test/internal/domain/model/profile"
	"github.com/xylong/bingo/test/internal/domain/model/user"
)

// FrontUserAgg 前台展示
type FrontUserAgg struct {
	User    *user.User       // 用户基础信息(聚合根)
	Profile *profile.Profile // 用户资料信息
}

func (u *FrontUserAgg) Get() {

}
