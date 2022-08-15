基于go语言gin框架的web项目骨架，专注于前后端分离业务场景

### 安装
```
go get -u github.com/xylong/bingo
```

### 功能
- [x] 控制器
- [x] 路由
    - [x] 分组
    - [ ] 限流
- [x] 中间件
    - [ ] 全局中间件
    - [x] 分组中间件
    - [ ] 路由中间件
- [ ] 日志
    - [ ] 日志分割
- [x] ioc容器
    - [x] 依赖注入

### 相关库
- [gin](https://github.com/gin-gonic/gin) http框架
- [logrus](https://github.com/sirupsen/logrus) 日志
- [gorm](gorm.io/gorm) orm
- [squirrel](github.com/Masterminds/squirrel) sql拼装工具

****

### 路由
路由支持GET、POST、PUT、PATCH、DELETE请求方式，并且支持通过Group方法对路由进行嵌套分组， bingo通过Route方法对路由函数进行注册。并且脚手架支持两种注册模式：
1. 控制器注册
2. Route方法注册
```go
// Route 路由注册函数
Route(group *bingo.Group) {
    group.POST("login", c.login)
    group.POST("register", c.smsRegister)
	
    group.Group(users *bingo.Group) {
        users.GET("users", c.index)
        users.GET("me", c.me)
        users.PUT("me", c.update)
        users.GET("users/:id", c.show)
        users.DELETE("logout", c.logout)
    }
}
```

****

### 控制器
1. 控制器：bingo.Controller接口，包含一个路由注册方法**Route(*bingo.Group)***，只需要实现该方法即是控制器。
2. 控制器方法：即路由函数，脚手架对gin.Context和gin.HandlerFunc进行了重载，支持直接返回结果值，函数签名为：
   1. **func(*bingo.Context) (int, string, interface{})***，该签名返回数据为json格式{"code":0,"message":"","data":null}
   2. **func(*bingo.Context) interface{}***，返回任意类型数据
   3. **func(*bingo.Context) string***，返回字符串
   4. **func(*bingo.Context)***，没有返回值
```go
// UserController 用户控制器
type UserController struct {}

// Route 注册路由
func (c *UserController) Route(group *bingo.Group) {
	group.POST("register", c.smsRegister)   // 可以在这里注册路由或者是外部注册路由
}

// register 注册
func (c *UserController) register(ctx *bingo.Context) interface{} {
	return c.service.Create(
		ctx.Binding(ctx.ShouldBind, &dto.SmsRegister{}).
			Unwrap().(*dto.SmsRegister))
}
```

****

### IOC与依赖注入
脚手架提供了一个轻量级的ioc容器，并且支持单例、多例的注入。
1. **Inject**方法，向容器传入依赖实体（***注：传入参数必须为结构体指针***），传入的实体如果包含有创建其它实体指针的方法将会被自动展开存入ioc容器当中
2. **inject**标签，实体字段通过标注该标签自动注入依赖
   1. 单例：`inject:"-"`，直接从ioc中获取依赖
   2. 多例：`inject:"结构体名.方法名"`，多例注入每次都将创建新的实体

#### 配置依赖管理
```go
// Service service依赖管理
type Service struct {
}

func NewService() *Service {
    return &Service{}
}

// User 创建UserService
func (c *Service) User() *service.UserService {
    return &service.UserService{
        Req: &assembler.UserReq{},
        Rep: &assembler.UserRep{},
    }
}
```
#### ioc
```go
// Service及UserService都会被保存在ioc中
bingo.Init().Inject(config.NewService()).Lunch()
```
#### 依赖注入
```go
type UserController struct {
    Service1 *service.UserService `inject:"-"`   // 单例注入
    Service2 *service.UserService `inject:"Service.User()"`   // 多例注入
}
```

****

### 使用
```go
func main() {
    bingo.Init().
        Mount("v1", v1.Controllers...)(middleware.NewLogger(), middleware.NewValidate()).
        Group("v2", func(group *bingo.Group) {
            group.POST("logoff", v2.Login)
        }, middleware.NewCsrf()).
        Lunch()
}
```
example在单元测试里

