package server

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	hServer "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/app/server"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/render"
	"github.com/spf13/viper"
)

func NewHttpServer(
	indexHlr server.IndexHandler,
) *hServer.Hertz {
	router := hServer.Default(hServer.WithHostPorts(fmt.Sprintf(":%d", viper.GetInt("http.port"))))
	// 测试URL
	router.GET("/ping", ping)

	// api的v1版本路由
	apiV1Router := router.Group("/api/v1")
	{
		// index测试方法
		apiV1Router.POST("/index", indexHlr.IndexV1)
	}
	return router
}

// 测试方法
func ping(c context.Context, ctx *app.RequestContext) {
	render.RespOK(ctx, "pong")
}
