package server

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/model/form"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/service"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/render"
)

type IndexHandler interface {
	IndexV1(_ context.Context, ctx *app.RequestContext)
}

func NewIndexHandler(
	indexSvc service.IndexService,
) IndexHandler {
	return &indexHandler{
		indexSvc: indexSvc,
	}
}

type indexHandler struct {
	indexSvc service.IndexService
}

func (ctr *indexHandler) IndexV1(_ context.Context, ctx *app.RequestContext) {
	var req form.IndexV1Req
	err := ctx.BindAndValidate(&req)
	if err != nil {
		render.RespError(ctx, 40001, fmt.Sprintf("%v", err))
		return
	}
	data := ctr.indexSvc.IndexV1()
	render.RespOK(ctx, data)
	return
}
