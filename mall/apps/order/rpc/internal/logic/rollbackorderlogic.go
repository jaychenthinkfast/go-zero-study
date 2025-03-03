package logic

import (
	"context"
	"mall/apps/order/rpc/internal/svc"
	"mall/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type RollbackOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRollbackOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RollbackOrderLogic {
	return &RollbackOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RollbackOrderLogic) RollbackOrder(in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	// todo: add your logic here and delete this line
	l.Infof("RollbackOrder: %+v\n", in)
	return &order.CreateOrderResponse{}, nil
}
