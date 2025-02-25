package mysql

import "github.com/cloudwego/biz-demo/go-mall/webapi/biz/model"

func CreateUsers(users []*model.User) error {
	return DB.Create(users).Error
}

func FindUserByNameOrEmail(userName, email string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := DB.Where("email = ?", email).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CheckUser(account, password string) ([]*model.User, error) {
	var res []*model.User
	if err := DB.Where("email = ?", account).Where("password = ?", password).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
