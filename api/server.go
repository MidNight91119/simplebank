package api

import (
	db "github.com/MidNight91119/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking services
type Server struct {
	store *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Add routes to router
	// if you pass multiple functions then the last one should be the handler
	// and in between should be middlewares
	router.POST("/accounts", server.createAccount)
	
	router.GET("/accounts/:id", server.getAccount)

	router.GET("/accounts", server.listAccounts)

	router.DELETE("/accounts/:id", server.deleteAccount)

	router.PUT("/accounts/:id", server.updateAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on specific address and listens
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}