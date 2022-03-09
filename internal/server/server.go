package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	r *gin.Engine
}

func New() *Server {
	r := NewRouter()
	return &Server{r: r}
}

func (s *Server) Start(addr string) error {
	return s.r.Run(addr)
}
