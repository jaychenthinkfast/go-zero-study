package logic

import (
	"context"

	"mall/apps/product/rpc/internal/svc"
	"mall/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLogic {
	return &ProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLogic) Product(in *product.ProductItemRequest) (*product.ProductItem, error) {
	// todo: add your logic here and delete this line
	p, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.ProductId)
	if err != nil {
		return nil, err
	}
	return &product.ProductItem{
		ProductId: p.Id,
		Name:      p.Name,
	}, nil
}
