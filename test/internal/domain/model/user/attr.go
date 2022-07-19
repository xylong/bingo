package user

import "github.com/xylong/bingo/test/internal/domain"

func WithID(id int) domain.Attr {
	return func(i interface{}) {
		if id > 0 {
			i.(*User).ID = id
		}
	}
}

func WithUnionid(unionid string) domain.Attr {
	return func(i interface{}) {
		if unionid != "" {
			i.(*User).WechatUnionid = unionid
		}
	}
}

func WithAppletOpenid(openid string) domain.Attr {
	return func(i interface{}) {
		if openid != "" {
			i.(*User).WechatAppletOpenid = openid
		}
	}
}

func WithOfficialOpenid(openid string) domain.Attr {
	return func(i interface{}) {
		if openid != "" {
			i.(*User).WechatOfficialOpenid = openid
		}
	}
}

func WithAvatar(avatar string) domain.Attr {
	return func(i interface{}) {
		if avatar != "" {
			i.(*User).Avatar = avatar
		}
	}
}

func WithNickName(name string) domain.Attr {
	return func(i interface{}) {
		if name != "" {
			i.(*User).Nickname = name
		}
	}
}

func WithPhone(phone string) domain.Attr {
	return func(i interface{}) {
		if phone != "" {
			i.(*User).Phone = phone
		}
	}
}

func WithEmail(email string) domain.Attr {
	return func(i interface{}) {
		if email != "" {
			i.(*User).Email = email
		}
	}
}
