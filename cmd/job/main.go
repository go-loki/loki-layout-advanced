package main

import (
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/config"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/log"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)
	logger.Info("start")

	app, cleanup, err := newApp(conf, logger)
	if err != nil {
		panic(err)
	}
	app.Run()
	defer cleanup()

}
