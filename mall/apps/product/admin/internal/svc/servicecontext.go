package svc

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mall/apps/product/admin/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	OssClient *oss.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	oc, err := oss.New(c.OSSEndpoint, c.AccessKeyID, c.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:    c,
		OssClient: oc,
	}
}
