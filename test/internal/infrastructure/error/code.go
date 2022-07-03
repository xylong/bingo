//go:generate stringer -type Code -linecomment
package error

// Code 错误吗
type Code int32

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
)
