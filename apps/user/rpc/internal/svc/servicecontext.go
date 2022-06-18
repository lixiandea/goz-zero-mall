package svc

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"lixiandea.github.com/go-zero-mall/apps/user/model"
	"lixiandea.github.com/go-zero-mall/apps/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UserModel
	ctx context.Context
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(conn, c.RedisCache),
	}
}
