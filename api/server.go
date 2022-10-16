package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tamarelhe/lets_game/api/token"
	db "github.com/tamarelhe/lets_game/db/sqlc"
	"github.com/tamarelhe/lets_game/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// user routes
	authRoutes.GET("/user/:id", server.getUser)
	authRoutes.GET("/users", server.getUsersList)
	authRoutes.DELETE("/user/:id", server.deleteUser)

	// group routes

	// game routes

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
