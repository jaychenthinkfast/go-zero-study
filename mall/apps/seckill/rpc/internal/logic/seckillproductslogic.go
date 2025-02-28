package logic

import (
	"context"

	"mall/apps/seckill/rpc/internal/svc"
	"mall/apps/seckill/rpc/seckill"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSeckillProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillProductsLogic {
	return &SeckillProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SeckillProductsLogic) SeckillProducts(in *seckill.SeckillProductsRequest) (*seckill.SeckillProductsResponse, error) {
	// todo: add your logic here and delete this line

	return &seckill.SeckillProductsResponse{}, nil
}
