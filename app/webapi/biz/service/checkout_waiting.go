package service

import (
	"context"
	"fmt"
	rpccheckout "github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/checkout"
	rpcpayment "github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/biz-demo/go-mall/webapi/hertz_gen/webapi/checkout"
	"github.com/cloudwego/biz-demo/go-mall/webapi/infra/rpc"
	"github.com/cloudwego/biz-demo/go-mall/webapi/mw"
	"github.com/cloudwego/hertz/pkg/app"
	utils2 "github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	userId := mw.GetUid(h.RequestContext)

	fmt.Print(userId)
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    userId,
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Address: &rpccheckout.Address{
			Country:       req.Country,
			ZipCode:       req.Zipcode,
			City:          req.City,
			State:         req.Province,
			StreetAddress: req.Street,
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
			CreditCardCvv:             req.Cvv,
		},
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return utils2.H{}, nil
}
