package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"mall/apps/product/rpc/internal/svc"
	"mall/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAndUpdateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAndUpdateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAndUpdateStockLogic {
	return &CheckAndUpdateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const (
	luaCheckAndUpdateScript = `
local counts = redis.call("HMGET", KEYS[1], "total", "seckill")
local total = tonumber(counts[1])
local seckill = tonumber(counts[2])
if seckill + 1 <= total then
	redis.call("HINCRBY", KEYS[1], "seckill", 1)
	return 1
end
return 0
`
)

func (l *CheckAndUpdateStockLogic) CheckAndUpdateStock(in *product.CheckAndUpdateStockRequest) (*product.CheckAndUpdateStockResponse, error) {
	// 这里需要前置条件，库存已经按stockKey规则存入 redis，这里在 makefile里硬编码写入 redis,
	// 实际可能需要定时脚本在商品预热时写入 redis
	val, err := l.svcCtx.BizRedis.EvalCtx(l.ctx, luaCheckAndUpdateScript, []string{stockKey(in.ProductId)})
	if err != nil {
		return nil, err
	}
	if val.(int64) == 0 {
		return nil, status.Errorf(codes.ResourceExhausted, fmt.Sprintf("insufficient stock: %d", in.ProductId))
	}
	return &product.CheckAndUpdateStockResponse{}, nil
}

func stockKey(pid uint64) string {
	return fmt.Sprintf("stock:%d", pid)
}
