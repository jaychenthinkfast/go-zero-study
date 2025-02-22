package logic

import (
	"context"

	"mall/apps/order/rpc/internal/svc"
	"mall/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrdersLogic {
	return &OrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrdersLogic) Orders(in *order.OrdersRequest) (*order.OrdersResponse, error) {
	// todo: add your logic here and delete this line
	uid := int64(123)
	if in.UserId == uid {
		orders := []*order.OrderItem{
			{
				OrderId:   "20220609123456",
				ProductId: 1,
				Quantity:  1,
			},
		}
		return &order.OrdersResponse{Orders: orders}, nil
	}
	return &order.OrdersResponse{}, nil
}
