package bingo

// Controller 控制器
type Controller interface {
	Route(group *Group)
}

// Bind 参数绑定
type Bind[T any] struct {
	try func(ctx *Context, t *T)
	catch func(ctx *Context, err error)
}

func NewBind[T any]() *Bind[T] {
	return &Bind[T]{}
}

// Try 参数绑定验证通过执行
func (b *Bind[T]) Try(f func(ctx *Context,t *T)) *Bind[T] {
	b.try = f
	return b
}

// Catch 失败执行
func (b *Bind[T]) Catch(f func(ctx *Context,err error)) *Bind[T] {
	b.catch=f
	return b
}

// Complete 完成调用
func (b *Bind[T]) Complete() func(ctx *Context) {
	return func(ctx *Context) {
		var t T

		if err:=ctx.ShouldBind(&t);err!=nil {
			b.catch(ctx,err)
		} else {
			b.try(ctx,&t)
		}
	}
}