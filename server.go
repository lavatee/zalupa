package zalupa

import (
	"net/http"
	"context"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(c context.Context) error {
	return s.httpServer.Shutdown(c)
}