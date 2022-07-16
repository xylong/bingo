package dto

import "time"

// 请求对象
type (
	// SmsRegister 短信注册
	SmsRegister struct {
		Phone    string `json:"phone" form:"phone" binding:"required"`        // 手机号
		Code     int32  `json:"code" form:"code" binding:"required"`          // 短信验证码
		Nickname string `json:"nickname" form:"nickname" binding:"omitempty"` // 昵称
	}

	// EmailRegister 邮箱注册
	EmailRegister struct {
		Email    string `json:"email" form:"email" binding:"required,email"`              // 邮箱
		Password string `json:"password" form:"password" binding:"required,min=6,max=18"` // 密码
		Nickname string `json:"nickname" form:"nickname" binding:"omitempty"`             // 昵称
	}

	// SimpleUserReq 简单用户请求参数
	SimpleUserReq struct {
		ID int `uri:"id" binding:"required,gt=0"`
	}
)

// 响应对象
type (
	// SimpleUser 简洁用户信息
	SimpleUser struct {
		ID       int    `json:"id"`
		Avatar   string `json:"avatar"`
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
	}

	// ExtraUser 额外用户信息
	ExtraUser struct {
		ID        int    `json:"id"`
		Phone     string `json:"phone"`
		Birthday  string `json:"birthday"`
		Gender    int    `json:"gender"`
		Level     int    `json:"level"`
		Signature string `json:"signature"`
		CreatedAt string `json:"created_at"`
	}

	// ThirdUser 第三方用户信息
	ThirdUser struct {
		ID                   int    `json:"id"`
		WechatUnionid        string `json:"wechat_unionid"`
		WechatAppletOpenid   string `json:"wechat_applet_openid"`
		WechatOfficialOpenid string `json:"wechat_official_openid"`
	}

	UserLog struct {
		ID   int       `json:"id"`
		Log  string    `json:"log"`
		Date time.Time `json:"date"`
	}

	UserInfo struct {
		ID       int    `json:"id"`
		Nickname string `json:"nickname"`
		Logs     []*UserLog
	}
)
