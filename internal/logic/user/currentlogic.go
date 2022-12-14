package user

import (
	"context"

	"go-zero-short/internal/svc"
	"go-zero-short/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurrentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentLogic {
	return &CurrentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurrentLogic) Current() (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line

	return &types.UserInfoResp{
		Id:          1,
		Name:        "候青晓",
		Mobile:      "",
		Avatar:      "",
		CreatedTime: "2022-08-21 15:12:00",
	}, nil
}
