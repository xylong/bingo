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
