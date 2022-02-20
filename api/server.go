package api

import (
	"github.com/NicolasMartino/simplebank/core"
	"github.com/gin-gonic/gin"
	_ "github.com/golang/mock/mockgen/model"
)

//Http server for our banking service
type Server struct {
	services *core.Services
	router   *gin.Engine
}

// Server constructor with dependency injection
func NewServer(services *core.Services) *Server {
	router := gin.Default()
	server := &Server{
		services: services,
	}
	// API routes
	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.GetAccount)
	router.GET("/accounts", server.GetAccountsWithPagination)
	router.PATCH("/accounts/:id", server.UpdateAccount)
	router.DELETE("/accounts/:id", server.DeleteAccount)

	server.router = router
	return server
}

// Start http server listenning on a specific adress
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
