package render

import (
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func RespOK(ctx *app.RequestContext, data any) {
	if data == nil {
		data = map[string]string{}
	}
	ctx.JSON(http.StatusOK, map[string]any{"data": data})
}

func RespError(ctx *app.RequestContext, code int, msg string) {
	ctx.JSON(http.StatusOK, map[string]any{"code": code, "msg": msg})
}

func RespEmpty(ctx *app.RequestContext, httpCode int) {
	ctx.String(httpCode, "")
}
