package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/biz-demo/go-mall/checkout/checkout_proto/infra/rpc"
	"github.com/cloudwego/biz-demo/go-mall/checkout/checkout_proto/kitex_gen/cart"
	"github.com/cloudwego/biz-demo/go-mall/checkout/checkout_proto/kitex_gen/checkout"
	"github.com/cloudwego/biz-demo/go-mall/checkout/checkout_proto/kitex_gen/order"
	"github.com/cloudwego/biz-demo/go-mall/checkout/checkout_proto/kitex_gen/payment"
	"github.com/cloudwego/biz-demo/go-mall/checkout/checkout_proto/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.

	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err)
		err = fmt.Errorf("GetCart.err:%v", err)
		return
	}
	if cartResult == nil || cartResult.Cart == nil || len(cartResult.Cart.Items) == 0 {
		err = errors.New("cart is empty")
		return
	}
	var (
		oi    []*order.OrderItem
		total float32
	)
	for _, cartItem := range cartResult.Cart.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if resultErr != nil {
			klog.Error(resultErr)
			err = resultErr
			return
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		cost := p.Price * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: cost,
		})
	}
	// create order
	orderReq := &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		OrderItems:   oi,
		Email:        req.Email,
	}
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(addr.ZipCode)
		orderReq.Address = &order.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       int32(zipCodeInt),
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		err = fmt.Errorf("PlaceOrder.err:%v", err)
		return
	}
	klog.Info("orderResult", orderResult)
	// empty cart
	emptyResult, err := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		err = fmt.Errorf("EmptyCart.err:%v", err)
		return
	}
	klog.Info(emptyResult)
	// charge
	var orderId string
	if orderResult != nil || orderResult.Order != nil {
		orderId = orderResult.Order.OrderId
	}
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		},
	}
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		err = fmt.Errorf("Charge.err:%v", err)
		return
	}

	klog.Info(paymentResult)
	// change order state
	klog.Info(orderResult)
	_, err = rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{UserId: req.UserId, OrderId: orderId})
	if err != nil {
		klog.Error(err)
		return
	}

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}
