package service

import (
	"context"
	rpccart "github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/product"
	common "github.com/cloudwego/biz-demo/go-mall/webapi/hertz_gen/webapi/common"
	"github.com/cloudwego/biz-demo/go-mall/webapi/infra/rpc"
	"github.com/cloudwego/biz-demo/go-mall/webapi/mw"
	"github.com/cloudwego/hertz/pkg/app"
	utils2 "github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	var items []map[string]string
	id := mw.GetUid(h.RequestContext)
	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: id,
	})
	if err != nil {
		return nil, err
	}
	var total float32
	for _, v := range carts.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: v.GetProductId()})
		if err != nil {
			continue
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{"Name": p.Name, "Description": p.Description, "Picture": p.Picture, "Price": strconv.FormatFloat(float64(p.Price), 'f', 2, 64), "Qty": strconv.Itoa(int(v.Quantity))})
		total += float32(v.Quantity) * p.Price
	}

	return utils2.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
