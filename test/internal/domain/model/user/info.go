package user

import "time"

// UserInfo 用户信息
type UserInfo struct {
	ID                   int       `gorm:"primaryKey;autoIncrement;" json:"id"`
	UserID               int       `gorm:"type:int(11);;not null;uniqueIndex;comment:用户🆔" json:"user_id"`
	WechatUnionid        string    `gorm:"type:varchar(100);default:null;uniqueIndex;comment:微信unionid" json:"wechat_unionid"`
	WechatAppletOpenid   string    `gorm:"type:varchar(100);default:null;uniqueIndex;comment:微信小程序🆔" json:"wechat_applet_openid"`
	WechatOfficialOpenid string    `gorm:"type:varchar(100);default:null;uniqueIndex;comment:微信公众号🆔" json:"wechat_official_openid"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func NewUserInfo() *UserInfo {
	return &UserInfo{}
}
