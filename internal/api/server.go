package api

import "github.com/gin-gonic/gin"

type Handler interface {
	Handler(context *gin.Context)
}

type Server struct {
	Port string
	Endpoints []Endpoint
}

type Endpoint struct {
	Route  string
	Method string
	Handle Handler
}

func (s *Server) Run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
