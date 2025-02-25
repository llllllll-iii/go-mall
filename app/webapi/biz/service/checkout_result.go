package service

import (
	"context"
	common "github.com/cloudwego/biz-demo/go-mall/webapi/hertz_gen/webapi/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutResultService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutResultService(Context context.Context, RequestContext *app.RequestContext) *CheckoutResultService {
	return &CheckoutResultService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutResultService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
