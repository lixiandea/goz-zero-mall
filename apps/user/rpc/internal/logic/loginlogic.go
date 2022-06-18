package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"lixiandea.github.com/go-zero-mall/apps/user/model"
	"lixiandea.github.com/go-zero-mall/apps/user/rpc/internal/svc"
	"lixiandea.github.com/go-zero-mall/apps/user/rpc/rpc"
	"lixiandea.github.com/go-zero-mall/common/cryptx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *rpc.LoginRequest) (*rpc.LoginResponse, error) {
	res, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil{
		if err == model.ErrNotFound{
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password !=res.Password{
		return nil, status.Error(100, "密码错误")
	}
	return &rpc.LoginResponse{
		ID: res.Id,
		Mobile: res.Mobile,
		Gender: res.Gender,
		Name: res.Mobile,
	}, nil
}
