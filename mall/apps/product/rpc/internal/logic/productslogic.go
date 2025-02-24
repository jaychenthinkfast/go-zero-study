package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/mr"
	"mall/apps/product/rpc/internal/model"
	"strconv"
	"strings"

	"mall/apps/product/rpc/internal/svc"
	"mall/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *product.ProductRequest) (*product.ProductResponse, error) {
	// todo: add your logic here and delete this line
	products := make(map[uint64]*product.ProductItem)
	pdis := strings.Split(in.ProductIds, ",")
	// 使用 mr.MapReduce 进行并行处理
	ps, err := mr.MapReduce(
		func(source chan<- any) {
			for _, pidStr := range pdis {
				pid, err := strconv.ParseUint(pidStr, 10, 64)
				if err != nil {
					l.Logger.Errorf("parse product id error: %v", err)
					continue
				}
				source <- pid
			}
		},
		func(item any, writer mr.Writer[any], cancel func(error)) {
			pid := item.(uint64)
			p, err := l.svcCtx.ProductModel.FindOne(l.ctx, pid)
			if err != nil {
				cancel(err)
				return
			}
			writer.Write(p)
		},
		func(pipe <-chan any, writer mr.Writer[any], cancel func(error)) {
			var r []*model.Product
			for p := range pipe {
				r = append(r, p.(*model.Product))
			}
			writer.Write(r)
		},
	)
	if err != nil {
		return nil, err
	}

	// 处理结果
	for _, p := range ps.([]*model.Product) {
		products[p.Id] = &product.ProductItem{
			ProductId: p.Id,
			Name:      p.Name,
		}
	}
	return &product.ProductResponse{Products: products}, nil
}
