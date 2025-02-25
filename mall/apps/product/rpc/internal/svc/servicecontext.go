package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/sync/singleflight"
	"mall/apps/product/rpc/internal/config"
	"mall/apps/product/rpc/internal/model"
)

type ServiceContext struct {
	Config        config.Config
	ProductModel  model.ProductModel
	CategoryModel model.CategoryModel
	BizRedis      *redis.Redis
	SingleGroup   singleflight.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:        c,
		ProductModel:  model.NewProductModel(conn, c.CacheRedis),
		CategoryModel: model.NewCategoryModel(conn, c.CacheRedis),
		BizRedis:      redis.MustNewRedis(c.BizRedis), //redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}
