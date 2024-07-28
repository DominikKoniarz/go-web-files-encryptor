package server

import "github.com/gin-gonic/gin"

type Server struct {
	Engine *gin.Engine
}

func New() *Server {
	return &Server{
		Engine: gin.Default(),
	}
}

func (s *Server) Run(addr string) {
	s.Engine.Run(addr)
}

func (s *Server) RegisterRoutes() {
	// Register routes here
	s.Engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

}
