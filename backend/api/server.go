// Server.go file
package api

import (
	db "github.com/awe8128/backend/db/sqlc"
	"github.com/awe8128/backend/util"
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
	// Set up the router to use JSON as the default content type
	router.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "up"})
	})

	// routes for the API
	router.POST("/createUser", server.createUser)
	router.GET("/login", server.login)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
