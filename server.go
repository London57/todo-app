package todo

import (
	"context"
	"net/http"

	"github.com/London57/todo-app/internal/usecase/auth"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	// s.httpServer = &http.Server{
	// 	Addr:         "127.0.0.1:" + port,
	// 	Handler:      handler,
	// 	ReadTimeout:  4 * time.Second,
	// 	WriteTimeout: 4 * time.Second,
	// }
	auth.RegisterUserByUsername()
	// return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
