package logic

import (
	"context"
	"mall/apps/order/rpc/internal/svc"
	"mall/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderCheckLogic {
	return &CreateOrderCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderCheckLogic) CreateOrderCheck(in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	// todo: add your logic here and delete this line
	l.Infof("CreateOrderCheck in:%v\n", in)
	return &order.CreateOrderResponse{}, nil
}
