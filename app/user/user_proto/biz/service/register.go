package service

import (
	"context"
	"errors"
	"github.com/cloudwego/biz-demo/go-mall/user/user_proto/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/go-mall/user/user_proto/biz/model"
	user "github.com/cloudwego/biz-demo/go-mall/user/user_proto/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	if req.Password != req.ConfirmPassword {
		err = errors.New("二次密码验证不一致")
		return
	}

	if err != nil {
		return
	}
	newUser := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}
	if err = model.Create(mysql.DB, s.ctx, newUser); err != nil {
		return
	}

	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
