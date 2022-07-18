// Code generated by "stringer -type Code -linecomment"; DO NOT EDIT.

package error

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OK-0]
	_ = x[ParamError-400]
	_ = x[Unauthorized-1001]
	_ = x[TokenMalformed-1002]
	_ = x[TokenSignatureInvalid-1003]
	_ = x[TokenExpired-1004]
	_ = x[TokenInvalid-1005]
	_ = x[ZeroIEntityD-10001]
	_ = x[NotFoundData-10002]
}

const (
	_Code_name_0 = "OK"
	_Code_name_1 = "ParamError"
	_Code_name_2 = "未授权令牌格式错误令牌签名验证失败令牌过期无效令牌"
	_Code_name_3 = "没有根实体🆔未找到数据"
)

var (
	_Code_index_2 = [...]uint8{0, 9, 27, 51, 63, 75}
	_Code_index_3 = [...]uint8{0, 19, 34}
)

func (i Code) String() string {
	switch {
	case i == 0:
		return _Code_name_0
	case i == 400:
		return _Code_name_1
	case 1001 <= i && i <= 1005:
		i -= 1001
		return _Code_name_2[_Code_index_2[i]:_Code_index_2[i+1]]
	case 10001 <= i && i <= 10002:
		i -= 10001
		return _Code_name_3[_Code_index_3[i]:_Code_index_3[i+1]]
	default:
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}