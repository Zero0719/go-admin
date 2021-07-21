package api

import (
	"go-admin/app/model"
	"go-admin/app/service"
	"go-admin/package/response"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var User = apiUser{}

type apiUser struct{}

// @Summary user-create
// @Description 新建用户
// @Param Authorization header string true "token"
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} response.ResponseStruct
// @Faulure 403,500 {object} response.ResponseStruct
// @Router /user [post]
func (*apiUser) Create(r *ghttp.Request) {
	var req *model.UserCreateRequest

	if err := r.Parse(&req); err != nil {
		response.Forbidden(r, gerror.Current(err).Error())
	}

	var serviceReq *model.UserServiceCreateRequest

	if err := gconv.Struct(req, &serviceReq); err != nil {
		response.Error(r, err.Error())
	}

	if err := service.User.Create(serviceReq); err != nil {
		response.Error(r, err.Error())
	}

	response.Success(r, "添加成功", nil)
}
