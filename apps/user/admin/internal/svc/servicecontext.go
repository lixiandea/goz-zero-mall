package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"lixiandea.github.com/go-zero-mall/apps/user/admin/internal/config"
	"lixiandea.github.com/go-zero-mall/apps/user/rpc/rpc"
)

type ServiceContext struct {
	Config config.Config
	UserRpc rpc.Rpc

}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRpc: rpc.NewRpc(zrpc.MustNewClient(c.UserRpc )) ,
	}
}
