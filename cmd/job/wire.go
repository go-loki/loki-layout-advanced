//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-hasaki/hasaki-layout-advanced/internal/job"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var JobSet = wire.NewSet(job.NewJob)

func newApp(*viper.Viper, *log.Logger) (*job.Job, func(), error) {
	panic(wire.Build(
		JobSet,
	))
}
