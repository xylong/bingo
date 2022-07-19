package domain

type (
	Attrs []Attr
	Attr  func(interface{})
)

func (a Attrs) Apply(v interface{}) {
	for _, attr := range a {
		attr(v)
	}
}
