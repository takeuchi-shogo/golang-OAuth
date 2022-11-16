package server

import (
	// "/Users/takeuchishougo/sampleapp/golang-OAuth/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	// Twitter *Twitter
	Gin  *gin.Engine
	Port string
}

// func NewServer(c *config.Config) *Server {
// 	server := &Server{
// 		Gin:  gin.Default(),
// 		Port: c.Port,
// 	}

// 	server.setRouting()
// 	return server
// }

func (s *Server) setRouting() {

	v1 := s.Gin.Group("/v1/api")
	{
		v1.GET("/authenticate/twitter", func(ctx *gin.Context) {})
	}
}

func (s *Server) Run(port string) {
	s.Gin.Run(port)
}
