package api

import (
	"github.com/NicolasMartino/simplebank/core"
	"github.com/NicolasMartino/simplebank/util"
	"github.com/gin-gonic/gin"
	_ "github.com/golang/mock/mockgen/model"
)

//Http server for our banking service
type Server struct {
	services  *core.Services
	appConfig *util.Config
	router    *gin.Engine
}

type URI struct {
	ID int64 `uri:"id" binding:"required"`
}

// Server constructor with dependency injection
func NewServer(services *core.Services, config *util.Config) *Server {
	router := gin.Default()
	server := &Server{
		services:  services,
		appConfig: config,
	}

	// API routes
	//Accounts
	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.GetAccount)
	router.GET("/accounts/user/:id", server.FindAccountsByUserId)
	router.GET("/accounts", server.GetAccountsWithPagination)
	router.PATCH("/accounts/:id", server.UpdateAccount)
	router.DELETE("/accounts/:id", server.DeleteAccount)

	//Transfer
	router.POST("/transfer", server.createTransfer)

	//User
	router.GET("/users/:id", server.FindUser)
	router.POST("/users", server.CreateUser)

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
