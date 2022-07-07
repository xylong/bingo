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
