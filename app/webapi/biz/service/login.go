package service

import (
	"context"
	"errors"
	"github.com/cloudwego/biz-demo/go-mall/webapi/biz/dal/mysql"
	auth "github.com/cloudwego/biz-demo/go-mall/webapi/hertz_gen/webapi/auth"
	"github.com/cloudwego/biz-demo/go-mall/webapi/utils"
	"github.com/cloudwego/hertz/pkg/app"
	utils2 "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func generateToken(user uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"identity": user,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("Lin123456789987654321"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (h *LoginService) Run(req *auth.LoginReq) (resp map[string]any, err error) {
	users, err := mysql.CheckUser(req.Email, utils.MD5(req.Password))
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {

		return nil, errors.New("账号不存在")
	}
	token, err := generateToken(users[0].ID)
	return utils2.H{
		"token": token,
	}, nil
}
