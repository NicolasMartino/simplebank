package api

import (
	db "simplt_bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

//Http server for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Server constructor with dependency injection
func NewServer(store *db.Store) *Server {
	router := gin.Default()
	server := &Server{
		store: store,
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
