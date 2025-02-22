package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mall/apps/app/api/internal/config"
	"mall/apps/order/rpc/orderclient"
	"mall/apps/product/rpc/productclient"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   orderclient.Order
	ProductRPC productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
