package user

// Info 用户信息
type Info struct {
	Avatar   string `gorm:"type:varchar(100);comment:头像" json:"avatar"`
	Nickname string `gorm:"type:varchar(20);not null;comment:昵称" json:"nickname"`
	Phone    string `gorm:"type:char(11);uniqueIndex;comment:手机号" json:"phone"`
	Email    string `gorm:"type:varchar(50);default:null;uniqueIndex;comment:邮件" json:"email"`
}

func NewInfo() *Info {
	return &Info{}
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
