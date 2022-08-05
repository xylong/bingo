package v2

import (
	"fmt"
	"github.com/xylong/bingo"
)

func init() {
	RegisterController(NewMockController())
}

type MockController struct{}

func NewMockController() *MockController {
	return &MockController{}
}

func (c *MockController) Foo(ctx *bingo.Context) {
	fmt.Println("测试定时器")
}

func (c *MockController) Route(group *bingo.Group) {

}
