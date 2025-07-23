package httpserver

import (
	"net"
	"strconv"
	"time"
)

type Option func(*Server)

func Address(host string, port int) Option {
	return func(s *Server) {
		s.address = net.JoinHostPort(host, strconv.Itoa(port))
	}
}

func Prefork(prefork bool) Option {
	return func(s *Server) {
		s.prefork = prefork
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.readTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.writeTimeout = timeout
	}
}

func ShutDownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
