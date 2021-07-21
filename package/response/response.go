package response

import (
	"net/http"

	"github.com/gogf/gf/net/ghttp"
)

type ResponseStruct struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(r *ghttp.Request, message string, data interface{}) {
	r.Response.WriteHeader(http.StatusOK)
	r.Response.WriteJson(&ResponseStruct{
		1, message, data,
	})
	r.Exit()
}

func Forbidden(r *ghttp.Request, message string) {
	r.Response.WriteHeader(http.StatusForbidden)
	r.Response.WriteJson(&ResponseStruct{
		0, message, nil,
	})
	r.Exit()
}

func UnAuthorized(r *ghttp.Request, message string) {
	r.Response.WriteHeader(http.StatusUnauthorized)
	r.Response.WriteJson(&ResponseStruct{
		0, message, nil,
	})
	r.Exit()
}

func Error(r *ghttp.Request, message string) {
	r.Response.WriteHeader(http.StatusInternalServerError)
	r.Response.WriteJson(&ResponseStruct{
		0, message, nil,
	})
	r.Exit()
}
