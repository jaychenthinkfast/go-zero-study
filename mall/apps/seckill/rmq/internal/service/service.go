package service

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"mall/apps/order/rpc/order"
	"mall/apps/order/rpc/orderclient"
	"mall/apps/product/rpc/product"
	"mall/apps/product/rpc/productclient"
	"mall/apps/seckill/rmq/internal/config"
	"sync"

	"github.com/zeromicro/go-zero/zrpc"
)

const (
	chanCount   = 10
	bufferCount = 1024
)

type Service struct {
	c          config.Config
	ProductRPC productclient.Product
	OrderRPC   orderclient.Order
	msgsChan   []chan *KafkaData
	waiter     sync.WaitGroup
}

type KafkaData struct {
	Uid uint64 `json:"uid"`
	Pid uint64 `json:"pid"`
}

func NewService(c config.Config) *Service {
	s := &Service{
		c:          c,
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		OrderRPC:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		msgsChan:   make([]chan *KafkaData, chanCount),
	}
	for i := 0; i < chanCount; i++ {
		ch := make(chan *KafkaData, bufferCount)
		s.msgsChan[i] = ch
		s.waiter.Add(1)
		go s.consume(ch)
	}

	return s
}

func (s *Service) consume(ch chan *KafkaData) {
	defer s.waiter.Done()

	for {
		m, ok := <-ch
		if !ok {
			log.Fatal("seckill rmq exit")
		}
		logx.Infof("consume msg: %+v\n", m)
		_, err := s.ProductRPC.CheckAndUpdateStock(context.Background(), &product.CheckAndUpdateStockRequest{ProductId: m.Pid})
		if err != nil {
			logx.Errorf("s.ProductRPC.CheckAndUpdateStock pid: %d error: %v", m.Pid, err)
			return
		}
		_, err = s.OrderRPC.CreateOrder(context.Background(), &order.CreateOrderRequest{Uid: m.Uid, Pid: m.Pid})
		if err != nil {
			logx.Errorf("CreateOrder uid: %d pid: %d error: %v", m.Uid, m.Pid, err)
			return
		}
		_, err = s.ProductRPC.UpdateProductStock(context.Background(), &product.UpdateProductStockRequest{ProductId: m.Pid, Num: 1})
		if err != nil {
			logx.Errorf("UpdateProductStock uid: %d pid: %d error: %v", m.Uid, m.Pid, err)
		}
	}
}

func (s *Service) Consume(_ context.Context, _ string, value string) error {
	logx.Infof("Consume value: %s\n", value)
	var data []*KafkaData
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}
	for _, d := range data {
		s.msgsChan[d.Pid%chanCount] <- d
	}
	return nil
}
