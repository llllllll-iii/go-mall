package mysql

import (
	"fmt"
	model "github.com/cloudwego/biz-demo/go-mall/order/order_proto/biz/model"
	"github.com/cloudwego/biz-demo/go-mall/order/order_proto/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	err := DB.AutoMigrate(&model.Order{}, &model.OrderItem{})
	if err != nil {
		fmt.Println("||||||", err.Error())
		return
	}
}
