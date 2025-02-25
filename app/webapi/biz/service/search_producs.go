package service

import (
	"context"
	rpcproduct "github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/biz-demo/go-mall/webapi/infra/rpc"
	utils2 "github.com/cloudwego/hertz/pkg/common/utils"

	product "github.com/cloudwego/biz-demo/go-mall/webapi/hertz_gen/webapi/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProducsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProducsService(Context context.Context, RequestContext *app.RequestContext) *SearchProducsService {
	return &SearchProducsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProducsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{Query: req.Q})
	if err != nil {
		return nil, err
	}
	return utils2.H{
		"items": p.Results,
		"q":     req.Q,
	}, nil
}
