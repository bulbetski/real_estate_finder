package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	r *gin.Engine
}

func New(token string) *Server {
	r := NewRouter(token)
	return &Server{r: r}
}

func (s *Server) Start(addr string) error {
	return s.r.Run(addr)
}

func (s *Server) LoadHTML(pattern string) {
	s.r.LoadHTMLGlob(pattern)
}
