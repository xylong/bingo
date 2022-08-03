package bingo

import (
	"bytes"
	"fmt"
	"github.com/xylong/bingo/utils"
	"regexp"
	"strings"
	"text/template"
)

const (
	VarPattern       = `[0-9a-zA-Z_]+`
	CompareSign      = ">|>=|<=|<|==|!="
	CompareSignToken = "gt|ge|le|lt|eq|ne" // gin模板支持的标识
	ComparePattern   = `^(` + VarPattern + `)\s*(` + CompareSign + `)\s*(` + VarPattern + `)\s*$`
)

type ComparableExpr string

func (e ComparableExpr) filter() string {
	reg, err := regexp.Compile(ComparePattern)
	if err != nil {
		return ""
	}

	result := reg.FindStringSubmatch(string(e))
	if result != nil && len(result) == 4 {
		if token := getCompareToken(result[2]); token != "" {
			return fmt.Sprintf("%s %s %s", token, parseToken(result[1]), parseToken(result[3]))
		}
	}

	return ""
}

// getCompareToken 根据比较符，获取token
func getCompareToken(sign string) string {
	for index, item := range strings.Split(CompareSign, "|") {
		if item == sign {
			return strings.Split(CompareSignToken, "|")[index]
		}
	}

	return ""
}

// parseToken 对于数字不加.(点)
func parseToken(token string) string {
	if utils.IsNumeric(token) {
		return token
	} else {
		return "." + token
	}
}

// IsComparableExpr 是否是"比较表达式"
func IsComparableExpr(expr string) bool {
	reg, err := regexp.Compile(ComparePattern)
	if err != nil {
		return false
	}
	return reg.MatchString(expr)
}

// ExecExpr 执行表达式，临时方法后期需要修改
func ExecExpr(expr string, data map[string]interface{}) (string, error) {
	tpl := template.New("expr").Funcs(map[string]interface{}{
		"echo": func(params ...interface{}) interface{} {
			return fmt.Sprintf("echo:%v", params[0])
		},
	})

	t, err := tpl.Parse(fmt.Sprintf("{{%s}}", ComparableExpr(expr).filter()))
	if err != nil {
		return "", err
	}

	var buf = &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
