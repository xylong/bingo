package user

import (
	"encoding/json"
	"github.com/xylong/bingo/test/internal/lib/db"
	"testing"
)

func init() {
	db.DB = db.InitGorm()
	db.DB.AutoMigrate(&User{}, &UserInfo{})
}

func TestNewUser(t *testing.T) {
	u := New(WithName("静静"), WithPhone("13666181506"), WithPassword("123456"),
		WithUnionid("sfjasfjijny373r3y472y233yr732y"), WithAppletOpenid("sajfiwqjqiwj38r92u02 92r93j2j0"))

	if err := db.DB.Create(u).Error; err != nil {
		t.Logf(err.Error())
	} else {
		j, _ := json.Marshal(u)
		t.Logf(string(j))
	}
}
