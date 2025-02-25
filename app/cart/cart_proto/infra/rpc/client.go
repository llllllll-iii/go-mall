package rpc

import (
	"github.com/cloudwego/biz-demo/go-mall/cart/cart_proto/conf"
	"github.com/cloudwego/biz-demo/go-mall/cart/cart_proto/kitex_gen/product/productcatalogservice"
	checkoututils "github.com/cloudwego/biz-demo/go-mall/cart/cart_proto/utils"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	err           error
	registryAddr  string
	serviceName   string
	commonSuite   client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		initProductClient()

	})
}

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

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: serviceName,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	checkoututils.MustHandleError(err)
}
