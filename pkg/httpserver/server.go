package httpserver

import (
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
	App *gin.Engine
}
