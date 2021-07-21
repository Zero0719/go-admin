package api

import (
	"go-admin/app/model"
	"go-admin/app/service"
	"go-admin/package/response"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Auth = apiAuth{}

type apiAuth struct{}

// @Summary Login
// @Description 后台登录
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} response.ResponseStruct
// @Faulure 403,500 {object} response.ResponseStruct
// @Router /login [post]
func (*apiAuth) Login(r *ghttp.Request) {
	var req *model.AuthLoginRequest

	if err := r.Parse(&req); err != nil {
		response.Forbidden(r, gerror.Current(err).Error())
	}

	token, err := service.Auth.Login(req)

	if  err != nil {
		response.Error(r, err.Error())
	}



	response.Success(r, "登录成功", g.Map{
		"token": token,
	})
}
