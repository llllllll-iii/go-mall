package service

import (
	"context"
	rpcproduct "github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/biz-demo/go-mall/webapi/infra/rpc"
	"github.com/cloudwego/hertz/pkg/common/utils"

	product "github.com/cloudwego/biz-demo/go-mall/webapi/hertz_gen/webapi/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: req.GetId()})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"item": p.Product,
	}, nil
}
