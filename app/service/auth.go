package service

import (
	"errors"
	"go-admin/app/model"
	"go-admin/package/jwt"

	"github.com/gogf/gf/crypto/gmd5"
)

var Auth = authService{}

type authService struct{}

func (*authService) Login(r *model.AuthLoginRequest) (token string, err error) {
	user := User.GetUserByUsername(r.Username)

	if user.Id <= 0 {
		return "", errors.New("用户不存在")
	}

	password, _ := gmd5.Encrypt(r.Password + user.Salt)

	if password != user.Password {
		return "", errors.New("密码不正确")
	}

	token, err = jwt.GenerateToken(user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}
