package dal

import (
	"github.com/cloudwego/biz-demo/go-mall/checkout/checkout_proto/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/go-mall/checkout/checkout_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
