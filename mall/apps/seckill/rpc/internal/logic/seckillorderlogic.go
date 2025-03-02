package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/limit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mall/apps/product/rpc/product"
	"mall/pkg/batcher"
	"strconv"
	"time"

	"mall/apps/seckill/rpc/internal/svc"
	"mall/apps/seckill/rpc/seckill"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	limitPeriod       = 10
	limitQuota        = 1
	seckillUserPrefix = "seckill#u#"
	localCacheExpire  = time.Second * 60

	batcherSize     = 100 //worker send queue size
	batcherBuffer   = 100 //worker channel buffer size
	batcherWorker   = 10
	batcherInterval = time.Second
)

type KafkaData struct {
	Uid uint64 `json:"uid"`
	Pid uint64 `json:"pid"`
}

type SeckillOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	limiter    *limit.PeriodLimit
	localCache *collection.Cache
	batcher    *batcher.Batcher
}

func NewSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillOrderLogic {
	localCache, err := collection.NewCache(localCacheExpire)
	if err != nil {
		panic(err)
	}
	s := &SeckillOrderLogic{
		ctx:        ctx,
		svcCtx:     svcCtx,
		Logger:     logx.WithContext(ctx),
		localCache: localCache,
		limiter:    limit.NewPeriodLimit(limitPeriod, limitQuota, svcCtx.BizRedis, seckillUserPrefix),
	}
	b := batcher.New(
		batcher.WithSize(batcherSize),
		batcher.WithBuffer(batcherBuffer),
		batcher.WithWorker(batcherWorker),
		batcher.WithInterval(batcherInterval),
	)
	b.Sharding = func(key string) int {
		pid, _ := strconv.ParseUint(key, 10, 64)
		return int(pid) % batcherWorker
	}
	b.Do = func(ctx context.Context, val map[string][]interface{}) {
		var msgs []*KafkaData
		for _, vs := range val {
			for _, v := range vs {
				msgs = append(msgs, v.(*KafkaData))
			}
		}
		kd, err := json.Marshal(msgs)
		if err != nil {
			s.Errorf("Batcher.Do json.Marshal msgs: %v error: %v", msgs, err)
		}
		if err = s.svcCtx.KafkaPusher.Push(ctx, string(kd)); err != nil {
			s.Errorf("KafkaPusher.Push kd: %s error: %v", string(kd), err)
		}
	}
	s.batcher = b
	s.batcher.Start()
	return s
}

func (l *SeckillOrderLogic) SeckillOrder(in *seckill.SeckillOrderRequest) (*seckill.SeckillOrderResponse, error) {
	// todo: add your logic here and delete this line
	code, _ := l.limiter.Take(strconv.FormatUint(in.UserId, 10))
	if code == limit.OverQuota {
		return nil, status.Errorf(codes.OutOfRange, "Number of requests exceeded the limit")
	}
	p, err := l.svcCtx.ProductRPC.Product(l.ctx, &product.ProductItemRequest{ProductId: in.ProductId})
	if err != nil {
		return nil, err
	}
	if p.Stock <= 0 {
		return nil, status.Errorf(codes.OutOfRange, "Insufficient stock")
	}
	if err = l.batcher.Add(strconv.FormatUint(in.ProductId, 10), &KafkaData{Uid: in.UserId, Pid: in.ProductId}); err != nil {
		l.Errorf("l.batcher.Add uid: %d pid: %d error: %v", in.UserId, in.ProductId, err)
	}
	return &seckill.SeckillOrderResponse{}, nil
}
