package transport

import (
	"go.uber.org/fx"
	"net/http"
)

type Route interface {
	http.Handler

	Pattern() string
}

func AsRoute(h any) any {
	return fx.Annotate(
		h,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
