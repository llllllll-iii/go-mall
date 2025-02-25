package cart

import (
	"context"
	"fmt"
	"github.com/cloudwego/biz-demo/go-mall/webapi/biz/service"
	"github.com/cloudwego/biz-demo/go-mall/webapi/biz/utils"
	cart "github.com/cloudwego/biz-demo/go-mall/webapi/hertz_gen/webapi/cart"
	common "github.com/cloudwego/biz-demo/go-mall/webapi/hertz_gen/webapi/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// AddCartItem .
// @router /cart [POST]
func AddCartItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddCartReq
	err = c.BindAndValidate(&req)
	fmt.Println(req.ProductNum, req.ProductId)
	if err != nil {
		fmt.Println(err.Error())
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	fmt.Println(req.ProductNum, req.ProductId)
	resp, err := service.NewAddCartItemService(ctx, c).Run(&req)
	if err != nil {
		fmt.Println(err.Error())
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	fmt.Println(req.ProductNum, req.ProductId)
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetCart .
// @router /cart [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	resp, err := service.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
