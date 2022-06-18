package svc

import (
	"lixiandea.github.com/go-zero-mall/apps/cart/admin/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
