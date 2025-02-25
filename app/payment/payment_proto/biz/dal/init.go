package dal

import (
	"github.com/cloudwego/biz-demo/go-mall/payment/payment_proto/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/go-mall/payment/payment_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
