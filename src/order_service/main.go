package main

import (
	"fmt"
	"monorepo/src/libs/log"
	"monorepo/src/order_service/configs"
	"monorepo/src/order_service/pkg/db"
	"monorepo/src/order_service/server"
	"monorepo/src/order_service/service"
	"monorepo/src/order_service/storage"
	"monorepo/src/order_service/tracer"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func loggerInit(config *configs.Configuration) log.Factory {
	fmt.Println("logInit")
	loggerForTracer, _ := zap.NewDevelopment(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)

	zapLogger := loggerForTracer.With(zap.String("service", config.ServiceName))
	logger := log.NewFactory(zapLogger)
	return logger
}
func main() {
	fmt.Println("main")

	app := fx.New(
		fx.Provide(
			configs.Config,
			db.Init,
			loggerInit,
			tracer.Load,
			storage.New,
			service.New,
		),

		fx.Invoke(
			server.Start,
		),

		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)

	app.Run()
}
