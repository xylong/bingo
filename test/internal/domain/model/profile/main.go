package profile

import (
	"fmt"
	"github.com/xylong/bingo/test/internal/domain/model"
	"gorm.io/gorm"
	"time"
)

const (
	Secrecy = -1 // 保密
	Female  = 0  // 女
	Male    = 1  // 男
)

// Profile 用户信息
type Profile struct {
	ID        int       `gorm:"primaryKey;autoIncrement;" json:"id"`
	UserID    int       `gorm:"type:int(11);;not null;uniqueIndex;comment:用户🆔" json:"user_id"`
	Password  string    `gorm:"type:varchar(32);comment:密码" json:"password"`
	Salt      string    `gorm:"type:char(6);comment:盐" json:"salt"`
	Birthday  time.Time `gorm:"type:date;default:null;comment:出生日期" json:"birthday"`
	Gender    int8      `gorm:"type:tinyint(1);default:-1;comment:-1保密 0女 1男" json:"gender"`
	Level     int8      `gorm:"type:tinyint(1);default:0;comment:等级" json:"level"`
	Signature string    `gorm:"type:varchar(255);comment=个性签名" json:"signature"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	*model.Model `gorm:"-"`
	//Repo repository.IProfileRepo `gorm:"-"`
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
	p.Filter(Age(20, model.Equal))
	return nil
}

func (p *Profile) Create() error {
	return nil
}

// Gender 根据性别比较
func Gender(gender int8, comparator int) model.Compare {
	return func(db *gorm.DB) *gorm.DB {
		switch comparator {
		case model.Equal:
			return db.Where("gender = ?", gender)
		default:
			return db
		}
	}
}

// Age 根据年龄比较
func Age(age int, comparator int) model.Compare {
	return func(db *gorm.DB) *gorm.DB {
		now := time.Now()
		year, month, day := now.Year(), now.Format("01"), now.Day()
		birthday := fmt.Sprintf("%d-%s-%d", year-age, month, day)

		switch comparator {
		case model.Equal:
			return db.Where("birthday >= ? and birthday <= ?", birthday, fmt.Sprintf("%d-12-31", year-age))
		case model.GreaterThan:
			return db.Where("birthday >= ?", birthday)
		case model.LessThan:
			return db.Where("birthday < ?", birthday)
		default:
			return db
		}
	}
}
