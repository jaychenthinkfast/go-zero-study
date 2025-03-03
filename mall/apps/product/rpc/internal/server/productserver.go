// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: product.proto

package server

import (
	"context"

	"mall/apps/product/rpc/internal/logic"
	"mall/apps/product/rpc/internal/svc"
	"mall/apps/product/rpc/product"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) Product(ctx context.Context, in *product.ProductItemRequest) (*product.ProductItem, error) {
	l := logic.NewProductLogic(ctx, s.svcCtx)
	return l.Product(in)
}

func (s *ProductServer) Products(ctx context.Context, in *product.ProductRequest) (*product.ProductResponse, error) {
	l := logic.NewProductsLogic(ctx, s.svcCtx)
	return l.Products(in)
}

func (s *ProductServer) ProductList(ctx context.Context, in *product.ProductListRequest) (*product.ProductListResponse, error) {
	l := logic.NewProductListLogic(ctx, s.svcCtx)
	return l.ProductList(in)
}

func (s *ProductServer) OperationProducts(ctx context.Context, in *product.OperationProductsRequest) (*product.OperationProductsResponse, error) {
	l := logic.NewOperationProductsLogic(ctx, s.svcCtx)
	return l.OperationProducts(in)
}

func (s *ProductServer) UpdateProductStock(ctx context.Context, in *product.UpdateProductStockRequest) (*product.UpdateProductStockResponse, error) {
	l := logic.NewUpdateProductStockLogic(ctx, s.svcCtx)
	return l.UpdateProductStock(in)
}

func (s *ProductServer) CheckAndUpdateStock(ctx context.Context, in *product.CheckAndUpdateStockRequest) (*product.CheckAndUpdateStockResponse, error) {
	l := logic.NewCheckAndUpdateStockLogic(ctx, s.svcCtx)
	return l.CheckAndUpdateStock(in)
}

func (s *ProductServer) CheckProductStock(ctx context.Context, in *product.UpdateProductStockRequest) (*product.UpdateProductStockResponse, error) {
	l := logic.NewCheckProductStockLogic(ctx, s.svcCtx)
	return l.CheckProductStock(in)
}

func (s *ProductServer) RollbackProductStock(ctx context.Context, in *product.UpdateProductStockRequest) (*product.UpdateProductStockResponse, error) {
	l := logic.NewRollbackProductStockLogic(ctx, s.svcCtx)
	return l.RollbackProductStock(in)
}
