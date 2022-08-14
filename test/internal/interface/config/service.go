package config

import (
	"github.com/xylong/bingo/test/internal/application/assembler"
	"github.com/xylong/bingo/test/internal/application/service"
)

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
