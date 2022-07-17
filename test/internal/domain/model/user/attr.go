package user

type Attr func(user *User)
type Attrs []Attr

func (a Attrs) apply(user *User) {
	for _, attr := range a {
		attr(user)
	}
}

func WithID(id int) Attr {
	return func(user *User) {
		if id > 0 {
			user.ID = id
		}
	}
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

func WithAvatar(avatar string) Attr {
	return func(user *User) {
		if avatar != "" {
			user.Avatar = avatar
		}
	}
}

func WithNickName(name string) Attr {
	return func(user *User) {
		if name != "" {
			user.Nickname = name
		}
	}
}

func WithPhone(phone string) Attr {
	return func(user *User) {
		if phone != "" {
			user.Phone = phone
		}
	}
}

func WithEmail(email string) Attr {
	return func(user *User) {
		if email != "" {
			user.Email = email
		}
	}
}
