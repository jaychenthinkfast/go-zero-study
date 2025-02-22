package logic

import (
	"context"

	"mall/apps/app/api/internal/svc"
	"mall/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 商品评论列表
func NewProductCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCommentLogic {
	return &ProductCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommentLogic) ProductComment(req *types.ProductCommentRequest) (resp *types.ProductCommentResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
