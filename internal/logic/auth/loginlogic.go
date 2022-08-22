package auth

import (
	"context"
	"go-zero-short/common/encrypt"
	"go-zero-short/common/errorx"
	"time"

	"go-zero-short/internal/svc"
	"go-zero-short/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	now := time.Now().Unix()
	token, err := encrypt.Token(l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire, now, 1)
	logx.Infof("token: %v", token)
	if err != nil {
		return nil, errorx.NewCodeError(10003, "用户登录失败")
	}
	return &types.LoginResp{
		Token:   token,
		Expires: now + l.svcCtx.Config.Auth.AccessExpire,
	}, nil
}
