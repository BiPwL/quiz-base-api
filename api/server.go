package api

import (
	"github.com/gin-gonic/gin"

	db "github.com/BiPwL/quiz-base-api/db/sqlc"
)

// Server serves HTTP requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// User handlers
	router.POST("users/new", server.createUser)
	router.GET("users/:id", server.getUser)
	router.GET("users", server.listUsers)
	router.DELETE("users/:id", server.deleteUser)
	router.POST("users", server.updateUser)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
