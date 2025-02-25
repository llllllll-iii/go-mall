package main

import (
	"context"
	"github.com/cloudwego/biz-demo/go-mall/checkout/checkout_proto/biz/service"
	checkout "github.com/cloudwego/biz-demo/go-mall/checkout/checkout_proto/kitex_gen/checkout"
)

// CheckoutServiceImpl implements the last service interface defined in the IDL.
type CheckoutServiceImpl struct{}

// Checkout implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	resp, err = service.NewCheckoutService(ctx).Run(req)

	return resp, err
}
