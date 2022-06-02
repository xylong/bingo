package bingo

// Middleware 中间件
type Middleware interface {
	Before(ctx *Context) error
}

type middlewares []Middleware

func (m middlewares) before(ctx *Context) {
	for _, middleware := range m {
		if err := middleware.Before(ctx); err != nil {
			panic(err)
		}
	}
}
