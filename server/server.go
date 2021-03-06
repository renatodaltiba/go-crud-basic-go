package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/renatodaltiba/rest-api-go/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   ":9000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Printf("Server running on port %s", s.port)
	log.Fatal(router.Run(s.port))
}
