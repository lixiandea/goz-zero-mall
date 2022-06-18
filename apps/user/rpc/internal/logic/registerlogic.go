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

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err == model.ErrNotFound{
		newUser := &model.User{
			Name:     in.Name,
			Mobile:   in.Mobile,
			Gender:   in.Gender,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password) ,
		}
		res, err := l.svcCtx.UserModel.Insert(l.ctx, newUser)
		if err != nil{
			return nil, status.Error(500, err.Error())
		}

		newUser.Id, err = res.LastInsertId()
		if err != nil{
			return nil, status.Error(500, err.Error())
		}
		return &rpc.RegisterResponse{
			ID: newUser.Id,
			Name: newUser.Name,
			Gender: newUser.Gender,
			Mobile: newUser.Mobile,
		}, nil
	}

	return nil , status.Error(500, err.Error())
}
