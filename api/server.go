package api

import (
	db "github.com/erncncbk/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server servis HTTP requests for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a spesific address
func (server *Server) Start(address string) error {
	return server.router.Run()
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
