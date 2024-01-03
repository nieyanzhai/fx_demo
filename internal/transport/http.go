package transport

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
	"net/http"
)

func NewHttpServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	server := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting http server", zap.String("addr", server.Addr))
			go server.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
	return server
}
