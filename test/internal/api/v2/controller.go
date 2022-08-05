package v2

import "github.com/xylong/bingo"

// Controllers 控制器
var Controllers = make([]bingo.Controller, 0)

func RegisterController(controller bingo.Controller) {
	Controllers = append(Controllers, controller)
}
