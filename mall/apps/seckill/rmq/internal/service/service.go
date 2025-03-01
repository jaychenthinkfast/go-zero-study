package service

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"mall/apps/order/rpc/order"
	"mall/apps/order/rpc/orderclient"
	"mall/apps/product/rpc/product"
	"mall/apps/product/rpc/productclient"
	"mall/apps/seckill/rmq/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type Service struct {
	c          config.Config
	ProductRPC productclient.Product
	OrderRPC   orderclient.Order
}

type KafkaData struct {
	Uid uint64 `json:"uid"`
	Pid uint64 `json:"pid"`
}

func NewService(c config.Config) *Service {
	return &Service{
		c:          c,
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		OrderRPC:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
	}
}

func (s *Service) Consume(ctx context.Context, _ string, value string) error {
	logx.Infof("Consume value: %s\n", value)
	var data KafkaData
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}
	p, err := s.ProductRPC.Product(ctx, &product.ProductItemRequest{ProductId: data.Pid})
	if err != nil {
		return err
	}
	if p.Stock <= 0 {
		return nil
	}
	_, err = s.OrderRPC.CreateOrder(ctx, &order.CreateOrderRequest{Uid: data.Uid, Pid: data.Pid})
	if err != nil {
		logx.Errorf("CreateOrder uid: %d pid: %d error: %v", data.Uid, data.Pid, err)
		return err
	}
	_, err = s.ProductRPC.UpdateProductStock(ctx, &product.UpdateProductStockRequest{ProductId: data.Pid, Num: 1})
	if err != nil {
		logx.Errorf("UpdateProductStock uid: %d pid: %d error: %v", data.Uid, data.Pid, err)
		return err
	}
	// TODO notify user of successful order placement
	return nil
}
