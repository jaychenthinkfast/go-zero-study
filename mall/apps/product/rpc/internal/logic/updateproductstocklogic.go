package logic

import (
	"context"

	"mall/apps/product/rpc/internal/svc"
	"mall/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductStockLogic {
	return &UpdateProductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductStockLogic) UpdateProductStock(in *product.UpdateProductStockRequest) (*product.UpdateProductStockResponse, error) {
	err := l.svcCtx.ProductModel.UpdateProductStock(l.ctx, in.ProductId, in.Num)
	if err != nil {
		return nil, err
	}
	return &product.UpdateProductStockResponse{}, nil
}
