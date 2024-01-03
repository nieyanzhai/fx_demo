package transport

import (
	"go.uber.org/zap"
	"io"
	"net/http"
)

type EchoHandler struct {
	log *zap.Logger
}

func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{
		log: log,
	}
}

func (h *EchoHandler) Pattern() string {
	return "/echo"
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.Warn("failed to copy body", zap.Error(err))
	}
}
