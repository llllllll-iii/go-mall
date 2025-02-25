package service

import (
	"context"
	rpcuser "github.com/cloudwego/biz-demo/go-mall/rpc_gen/kitex_gen/user"
	auth "github.com/cloudwego/biz-demo/go-mall/webapi/hertz_gen/webapi/auth"
	"github.com/cloudwego/biz-demo/go-mall/webapi/infra/rpc"
	utils3 "github.com/cloudwego/biz-demo/go-mall/webapi/utils"
	"github.com/cloudwego/hertz/pkg/app"
	utils2 "github.com/cloudwego/hertz/pkg/common/utils"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp map[string]any, err error) {
	_, err = rpc.UserClient.Register(h.Context, &rpcuser.RegisterReq{
		Email:           req.Email,
		Password:        utils3.MD5(req.Password),
		ConfirmPassword: utils3.MD5(req.PasswordConfirm),
	})
	if err != nil {
		return nil, err
	}

	return utils2.H{
		"msg": "注册成功",
	}, nil
}
