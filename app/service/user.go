package service

import (
	"errors"
	"go-admin/app/dao"
	"go-admin/app/model"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/util/grand"
)

var User = userService{}

type userService struct{}

func (u *userService) Create(r *model.UserServiceCreateRequest) error {
	user := u.GetUserByUsername(r.Username)

	if user.Id > 0 {
		return errors.New("用户名已存在")
	}

	// salt 随机生成
	r.Salt = grand.Letters(20)
	r.Status = 1
	r.Password, _ = gmd5.Encrypt(r.Password+r.Salt)

	if _, err := dao.User.Save(r); err != nil {
		return err
	}

	return nil
}

func (u *userService) GetUserByUsername(username string) model.User {
	var user model.User
	dao.User.Where("username", username).Scan(&user)
	return user
}
