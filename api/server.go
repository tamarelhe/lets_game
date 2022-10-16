package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/tamarelhe/lets_game/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// user routes
	router.POST("/user", server.createUser)
	router.GET("/user/:id", server.getUser)
	router.GET("/users", server.getUsersList)
	router.DELETE("/user/:id", server.deleteUser)

	// group routes

	// game routes

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
