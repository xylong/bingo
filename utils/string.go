package utils

import "fmt"

// ZeroFill 字符串补0
// str 要操作的字符串;length 结果字符串长度;flag true前置补0，false后置补0
func ZeroFill(str string, length int, flag bool) string {
	if len(str) > length || length <= 0 {
		return str
	}

	if flag {
		return fmt.Sprintf("%0*s", length, str) // 不足前置补零
	}

	result, l := str, len(str)
	for i := 0; i < l; i++ {
		result += "0"
	}

	return result
}
