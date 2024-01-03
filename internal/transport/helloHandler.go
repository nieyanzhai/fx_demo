package transport

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type HelloHandler struct {
	log *zap.Logger
}

func NewHelloHandler(log *zap.Logger) *HelloHandler {
	return &HelloHandler{
		log: log,
	}
}

func (h *HelloHandler) Pattern() string {
	return "/hello"
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("failed to read body", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "Hello, %s!", body)
	if err != nil {
		h.log.Error("failed to write body", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
