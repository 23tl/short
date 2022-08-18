package testing

import (
	"context"
	"time"

	"go-zero-short/internal/svc"
	"go-zero-short/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestingLogic {
	return &TestingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestingLogic) Testing() (resp *types.TestResp, err error) {
	dateNow := time.Now()
	return &types.TestResp{
		Date: dateNow.Format("2006-01-02 15:04:05"),
		Name: l.svcCtx.Config.Name,
	}, nil
}
