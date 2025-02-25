package dal

import (
	"github.com/cloudwego/biz-demo/go-mall/webapi/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/go-mall/webapi/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
