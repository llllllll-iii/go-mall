package rpc

import (
	"github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/biz-demo/go-mall/webapi/utils"

	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"

	"github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/biz-demo/go-mall/webapi/conf"
	"sync"

	"github.com/cloudwego/kitex/client"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type CommonGrpcClientSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonGrpcClientSuite) Options() []client.Option {
	r, err := consul.NewConsulResolver(s.RegistryAddr)
	if err != nil {
		panic(err)
	}
	opts := []client.Option{
		client.WithResolver(r),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
	}

	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		client.WithSuite(tracing.NewClientSuite()),
	)

	return opts
}

var (
	CheckoutClient checkoutservice.Client
	CartClient     cartservice.Client
	ProductClient  productcatalogservice.Client
	PaymentClient  paymentservice.Client
	OrderClient    orderservice.Client
	UserClient     userservice.Client
	once           sync.Once
	err            error
	registryAddr   string
	commonSuite    client.Option
)

const ServiceName = "webapi"

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Hertz.RegistryAddress
		commonSuite = client.WithSuite(CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: ServiceName,
		})
		initUserClient()
		initProductClient()
		initCartClient()
		initPaymentClient()
		initOrderClient()
		initCheckoutClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	utils.MustHandleError(err)
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
	utils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	utils.MustHandleError(err)
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", commonSuite)
	utils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	utils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", commonSuite)
	utils.MustHandleError(err)
}
