package config

import (
	"github.com/xylong/bingo/test/internal/application/assembler"
	"github.com/xylong/bingo/test/internal/application/service"
)

type UserServiceConfig struct {
}

func NewUserServiceConfig() *UserServiceConfig {
	return &UserServiceConfig{}
}

func (c *UserServiceConfig) UserService() *service.UserService {
	return &service.UserService{
		Req: &assembler.UserReq{},
		Rep: &assembler.UserRep{},
	}
}
