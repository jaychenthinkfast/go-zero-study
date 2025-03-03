package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/dtm-labs/dtmgrpc"
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
		//go s.consume(ch)
		go s.consumeDTM(ch)
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

var dtmServer = "etcd://localhost:2379/dtmservice"

func (s *Service) consumeDTM(ch chan *KafkaData) {
	defer s.waiter.Done()

	productServer, err := s.c.ProductRPC.BuildTarget()
	if err != nil {
		log.Fatalf("s.c.ProductRPC.BuildTarget error: %v", err)
	}
	orderServer, err := s.c.OrderRPC.BuildTarget()
	if err != nil {
		log.Fatalf("s.c.OrderRPC.BuildTarget error: %v", err)
	}

	for {
		m, ok := <-ch
		if !ok {
			log.Fatal("seckill rmq exit")
		}
		fmt.Printf("consume msg: %+v\n", m)

		gid := dtmgrpc.MustGenGid(dtmServer)
		err := dtmgrpc.TccGlobalTransaction(dtmServer, gid, func(tcc *dtmgrpc.TccGrpc) error {
			if e := tcc.CallBranch(
				&product.UpdateProductStockRequest{ProductId: m.Pid, Num: 1},
				productServer+"/product.Product/CheckProductStock",
				productServer+"/product.Product/UpdateProductStock",
				productServer+"/product.Product/RollbackProductStock",
				&product.UpdateProductStockRequest{}); err != nil {
				logx.Errorf("tcc.CallBranch server: %s error: %v", productServer, err)
				return e
			}
			if e := tcc.CallBranch(
				&order.CreateOrderRequest{Uid: m.Uid, Pid: m.Pid},
				orderServer+"/order.Order/CreateOrderCheck",
				orderServer+"/order.Order/CreateOrder",
				orderServer+"/order.Order/RollbackOrder",
				&order.CreateOrderResponse{},
			); err != nil {
				logx.Errorf("tcc.CallBranch server: %s error: %v", orderServer, err)
				return e
			}
			return nil
		})
		logger.FatalIfError(err)
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
