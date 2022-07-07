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
