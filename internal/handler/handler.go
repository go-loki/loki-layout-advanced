package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/middleware"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/log"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}
func GetUserIdFromCtx(ctx *app.RequestContext) string {
	v, exists := ctx.Get("claims")
	if !exists {
		return ""
	}
	return v.(*middleware.MyCustomClaims).UserId
}
