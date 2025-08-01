package httpserver

import (
	"context"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	_defautlAddr            = ":80"
	_defaultReadTimout      = 5 * time.Second
	_defaultWriteTimout     = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	App        *gin.Engine
	HTTPServer *http.Server
	notify     chan error

	address         string
	prefork         bool
	readTimeout     time.Duration
	writeTimeout    time.Duration
	shutdownTimeout time.Duration
}

func New(opts ...Option) *Server {
	s := &Server{
		App:             nil,
		HTTPServer:      nil,
		notify:          make(chan error, 1),
		address:         _defautlAddr,
		readTimeout:     _defaultReadTimout,
		writeTimeout:    _defaultWriteTimout,
		shutdownTimeout: _defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(s)
	}
	httpserv := &http.Server{
		Addr:         s.address,
		WriteTimeout: s.writeTimeout,
		ReadTimeout:  s.readTimeout,
	}
	s.HTTPServer = httpserv
	g := gin.New()
	if s.prefork {
		runtime.GOMAXPROCS(1)
	}
	s.App = g
	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.HTTPServer.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	context, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.HTTPServer.Shutdown(context)
}
