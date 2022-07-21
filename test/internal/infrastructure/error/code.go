package error

//go:generate stringer -type Code -linecomment
type Code uint16

const (
	OK         Code = 0
	ParamError Code = 400

	Unauthorized          Code = 1001 // 未授权
	TokenMalformed        Code = 1002 // 令牌格式错误
	TokenSignatureInvalid Code = 1003 // 令牌签名验证失败
	TokenExpired          Code = 1004 // 令牌过期
	TokenInvalid          Code = 1005 // 无效令牌

	ZeroIEntityD Code = 10001 // 没有实体🆔
	NotFoundData Code = 10002 // 未找到数据
	InsertError  Code = 10003 // 数据创建错误

	CreateUserError    Code = 20101 // 用户创建失败
	CreateProfileError Code = 20102 // 用户信息信息创建失败
	CreateUserLogError Code = 20103 // 用户日志创建失败
)
