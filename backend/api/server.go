// Server.go file
package api

import (
	"github.com/AlexInJapan/alex-app/backend/util"
	db "github.com/AlexInJapan/alex-app/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP request
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{store: store}
	router := gin.Default()

	// routes for the API
	router.POST("/user", server.createUser)
	router.GET("/get_user", server.getUser)
	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
