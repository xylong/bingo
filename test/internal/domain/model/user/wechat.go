package user

// Wechat 微信
type Wechat struct {
	WechatUnionid        string `gorm:"type:varchar(100);default:null;uniqueIndex;comment:微信unionid" json:"wechat_unionid"`
	WechatAppletOpenid   string `gorm:"type:varchar(100);default:null;uniqueIndex;comment:微信小程序🆔" json:"wechat_applet_openid"`
	WechatOfficialOpenid string `gorm:"type:varchar(100);default:null;uniqueIndex;comment:微信公众号🆔" json:"wechat_official_openid"`
}

func NewWechat() *Wechat {
	return &Wechat{}
}

func WithUnionid(unionid string) Attr {
	return func(user *User) {
		if unionid != "" {
			user.WechatUnionid = unionid
		}
	}
}

func WithAppletOpenid(openid string) Attr {
	return func(user *User) {
		if openid != "" {
			user.WechatAppletOpenid = openid
		}
	}
}

func WithOfficialOpenid(openid string) Attr {
	return func(user *User) {
		if openid != "" {
			user.WechatOfficialOpenid = openid
		}
	}
}
