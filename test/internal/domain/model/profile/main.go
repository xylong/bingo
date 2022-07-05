package profile

import (
	"github.com/xylong/bingo/test/internal/domain/repository"
	"time"
)

type (
	Attrs []Attr
	Attr  func(profile *Profile)
)

func (a Attrs) apply(profile *Profile) {
	for _, attr := range a {
		attr(profile)
	}
}

func WithUserID(id int) Attr {
	return func(profile *Profile) {
		if id > 0 {
			profile.UserID = id
		}
	}
}

func WithPassword(password string) Attr {
	return func(profile *Profile) {
		if len(password) > 0 {
			profile.Password = password
		}
	}
}

func WithBirthday(birthday string) Attr {
	return func(profile *Profile) {
		if len(birthday) > 0 {
			if t, err := time.Parse("2006-01-02 15:04:05", birthday); err == nil {
				profile.Birthday = t
			}
		}
	}
}

func WithGender(gender int) Attr {
	return func(profile *Profile) {
		profile.Gender = gender
	}
}

func WithLevel(level int) Attr {
	return func(profile *Profile) {
		if level >= 0 {
			profile.Level = level
		}
	}
}

func WithSignature(signature string) Attr {
	return func(profile *Profile) {
		if signature != "" {
			profile.Signature = signature
		}
	}
}

func WithRepo(repo repository.IProfileRepo) Attr {
	return func(profile *Profile) {
		profile.Repo = repo
	}
}

// Profile 用户信息
type Profile struct {
	ID        int       `gorm:"primaryKey;autoIncrement;" json:"id"`
	UserID    int       `gorm:"type:int(11);;not null;uniqueIndex;comment:用户🆔" json:"user_id"`
	Password  string    `gorm:"type:varchar(32);comment:密码" json:"password"`
	Birthday  time.Time `gorm:"type:date;default:null;comment:出生日期" json:"birthday"`
	Gender    int       `gorm:"type:tinyint(1);default:-1;comment:-1保密 0女 1男" json:"gender"`
	Level     int       `gorm:"type:tinyint(1);default:0;comment:等级" json:"level"`
	Signature string    `goem:"type:varchar(255);comment=个性签名" json:"signature"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Repo repository.IProfileRepo `gorm:"-"`
}

func New(attr ...Attr) *Profile {
	p := &Profile{}
	Attrs(attr).apply(p)

	return p
}

func (p *Profile) Name() string {
	return "profile"
}

func (p *Profile) Get() error {
	return p.Repo.GetByUser(p)
}

func (p *Profile) Create() error {
	return p.Repo.Create(p)
}
