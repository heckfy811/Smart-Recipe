package api

import (
	"github.com/gin-gonic/gin"
)

func (s *APIServer) Run() {
	app := gin.Default()
	app.Run(s.Addr)
}

type APIServer struct {
	Addr string
}
