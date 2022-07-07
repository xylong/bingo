package profile

import (
	"github.com/xylong/bingo/test/internal/domain/repository"
	"time"
)

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
