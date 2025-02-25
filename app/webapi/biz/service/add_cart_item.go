package service

import (
	"context"
	"fmt"
	rpccart "github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/cart"
	cart "github.com/cloudwego/biz-demo/go-mall/webapi/hertz_gen/webapi/cart"
	common "github.com/cloudwego/biz-demo/go-mall/webapi/hertz_gen/webapi/common"
	"github.com/cloudwego/biz-demo/go-mall/webapi/infra/rpc"
	"github.com/cloudwego/biz-demo/go-mall/webapi/mw"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp *common.Empty, err error) {
	fmt.Println(12121)

	uid := mw.GetUid(h.RequestContext)

	fmt.Println(uid)
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: uid,
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.ProductNum,
		},
	})

	fmt.Println(err)
	fmt.Println(12121)
	return nil, err
}
