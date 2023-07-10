//go:build wireinject
// +build wireinject

package main

import (
	hserver "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/handler"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/middleware"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/repository"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/server"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/service"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/helper/sid"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(server.NewServer)

var SidSet = wire.NewSet(sid.NewSid)

var JwtSet = wire.NewSet(middleware.NewJwt)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)

func newApp(*viper.Viper, *log.Logger) (*hserver.Hertz, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
		SidSet,
		JwtSet,
	))
}
