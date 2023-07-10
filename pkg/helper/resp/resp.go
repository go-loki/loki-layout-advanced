package resp

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx *app.RequestContext, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: 0, Message: "success", Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *app.RequestContext, httpCode, code int, message string, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: code, Message: message, Data: data}
	ctx.JSON(httpCode, resp)
}
