package logic

import (
	"context"

	"inet/microserver/gozero/httpdemo/internal/svc"
	"inet/microserver/gozero/httpdemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HttpdemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHttpdemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HttpdemoLogic {
	return &HttpdemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HttpdemoLogic) Httpdemo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
