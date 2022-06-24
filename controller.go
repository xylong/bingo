package bingo

// Controller 控制器
type Controller interface {
	Route(group *Group)
}

// Bind 参数绑定
type Bind[T any] struct {
	err   error
	try   func(ctx *Context, t *T) any
	catch func(ctx *Context, err error)
}

func NewBind[T any]() *Bind[T] {
	return &Bind[T]{}
}

// Try 参数绑定验证通过执行
func (b *Bind[T]) Try(f func(ctx *Context, t *T) any) *Bind[T] {
	b.try = f
	return b
}

// Catch 失败执行
func (b *Bind[T]) Catch(f ...func(ctx *Context, err error)) *Bind[T] {
	if len(f) > 0 {
		b.catch = f[0]
	} else {
		b.catch = func(ctx *Context, err error) {
			b.err = err
		}
	}

	return b
}

// Complete 完成调用
func (b *Bind[T]) Complete() func(ctx *Context) any {
	return func(ctx *Context) any {
		var t T

		if err := ctx.ShouldBind(&t); err != nil {
			b.catch(ctx, err)
			return b.err
		} else {
			return b.try(ctx, &t)
		}
	}
}
