package model

type AuthLoginRequest struct {
	Username string `p:"username" v:"required|length:6,20#用户名不能为空|用户名长度在:min到:max之间"`
	Password string `p:"password" v:"required|length:6,20#密码不能为空|密码长度在:min到:max之间"`
}
