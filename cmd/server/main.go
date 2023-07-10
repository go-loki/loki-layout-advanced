package main

import (
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/config"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/log"
	"go.uber.org/zap"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	app, cleanup, err := newApp(conf, logger)
	if err != nil {
		panic(err)
	}
	logger.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))

	app.Spin()
	defer cleanup()

}
