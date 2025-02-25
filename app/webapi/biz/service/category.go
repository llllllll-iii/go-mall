package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/biz-demo/go-mall/webapi/infra/rpc"
	utils2 "github.com/cloudwego/hertz/pkg/common/utils"

	category "github.com/cloudwego/biz-demo/go-mall/webapi/hertz_gen/webapi/category"
	"github.com/cloudwego/hertz/pkg/app"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	fmt.Println(req.Category)
	p, _ := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{CategoryName: req.Category})
	return utils2.H{
		"title": "Category",
		"items": p.Products,
	}, nil
}
