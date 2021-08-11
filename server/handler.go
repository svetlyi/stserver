package server

import (
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type Server struct {
	basicFileServer http.Handler
	logger          *zap.SugaredLogger
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Infof(
		"request from %s: %s; user-agent: %v; url: %s",
		r.RemoteAddr,
		r.URL.String(),
		r.Header.Get("User-Agent"),
		r.URL.String(),
	)
	if !strings.HasSuffix(r.URL.String(), "/") && r.URL.String() != "" {
		w.Header().Set("Content-Disposition", "attachment; filename="+r.URL.String())
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	}

	s.basicFileServer.ServeHTTP(w, r)
}

func NewFileServerHandler(root string, logger *zap.SugaredLogger) http.Handler {
	handler := http.FileServer(http.Dir(root))

	return Server{basicFileServer: handler, logger: logger}
}
