package main

import (
	"fx_demo/internal/transport"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(
			transport.NewHttpServer,
			fx.Annotate(
				transport.NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			transport.AsRoute(transport.NewEchoHandler),
			transport.AsRoute(transport.NewHelloHandler),
			zap.NewExample,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
