package home

import (
	"context"

	"github.com/cloudwego/biz-demo/go-mall/webapi/biz/service"
	"github.com/cloudwego/biz-demo/go-mall/webapi/biz/utils"
	home "github.com/cloudwego/biz-demo/go-mall/webapi/hertz_gen/webapi/home"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Method1 .
// @router / [GET]
func Method1(ctx context.Context, c *app.RequestContext) {
	var err error
	var req home.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &home.Empty{}
	resp, err = service.NewMethod1Service(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
