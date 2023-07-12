package server

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/handler"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/middleware"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/helper/resp"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/log"
	"github.com/spf13/viper"
)

func NewServer(
	conf *viper.Viper,
	logger *log.Logger,
	jwt *middleware.JWT,
	userHandler handler.UserHandler,
) *server.Hertz {
	r := server.Default(server.WithHostPorts(fmt.Sprintf(":%d", conf.GetInt("http.port"))))

	r.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	// 无权限路由
	noAuthRouter := r.Group("/").Use(middleware.RequestLogMiddleware(logger))
	{

		noAuthRouter.GET("/", func(c context.Context, ctx *app.RequestContext) {
			logger.WithContext(ctx).Info("hello")
			resp.HandleSuccess(ctx, map[string]interface{}{
				"say": "Hi Hasaki!",
			})
		})

		noAuthRouter.POST("/register", userHandler.Register)
		noAuthRouter.POST("/login", userHandler.Login)
	}
	// 非严格权限路由
	noStrictAuthRouter := r.Group("/").Use(middleware.NoStrictAuth(jwt, logger), middleware.RequestLogMiddleware(logger))
	{
		noStrictAuthRouter.GET("/user", userHandler.GetProfile)
	}

	// 严格权限路由
	strictAuthRouter := r.Group("/").Use(middleware.StrictAuth(jwt, logger), middleware.RequestLogMiddleware(logger))
	{
		strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
	}

	return r
}
