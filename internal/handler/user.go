package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/service"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/helper/resp"
	"github.com/pkg/errors"
	"net/http"
)

type UserHandler interface {
	Register(_ context.Context, ctx *app.RequestContext)
	Login(_ context.Context, ctx *app.RequestContext)
	GetProfile(_ context.Context, ctx *app.RequestContext)
	UpdateProfile(_ context.Context, ctx *app.RequestContext)
}

type userHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) UserHandler {
	return &userHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (h *userHandler) Register(_ context.Context, ctx *app.RequestContext) {
	req := new(service.RegisterRequest)
	if err := ctx.BindAndValidate(req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	if err := h.userService.Register(ctx, req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, nil)
}

func (h *userHandler) Login(_ context.Context, ctx *app.RequestContext) {
	var req service.LoginRequest
	if err := ctx.BindAndValidate(&req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	token, err := h.userService.Login(ctx, &req)
	if err != nil {
		resp.HandleError(ctx, http.StatusUnauthorized, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, utils.H{
		"accessToken": token,
	})
}

func (h *userHandler) GetProfile(_ context.Context, ctx *app.RequestContext) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		resp.HandleError(ctx, http.StatusUnauthorized, 1, "unauthorized", nil)
		return
	}

	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, user)
}

func (h *userHandler) UpdateProfile(_ context.Context, ctx *app.RequestContext) {
	userId := GetUserIdFromCtx(ctx)

	var req service.UpdateProfileRequest
	if err := ctx.BindAndValidate(&req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, nil)
}
