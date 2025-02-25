package mysql

import (
	"fmt"
	"github.com/cloudwego/biz-demo/go-mall/user/user_proto/biz/model"
	"github.com/cloudwego/biz-demo/go-mall/user/user_proto/conf"
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
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("||||||", err.Error())
		return
	}
}
