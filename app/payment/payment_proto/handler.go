package main

import (
	"context"
	payment "github.com/cloudwego/biz-demo/go-mall/payment/payment_proto/kitex_gen/payment"
	"github.com/cloudwego/biz-demo/go-mall/payment/payment_proto/biz/service"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)

	return resp, err
}
