package api

import (
	"log"
	"net/http"

	db "github.com/awe8128/backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Name:  req.Name,
		Email: req.Email,
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, user)

}

type createGetUserRequest struct {
	Email string `form:"email" binding:"required"`
}

func (server *Server) login(ctx *gin.Context) {
	var req createGetUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Println("email", req.Email)

	// user, err := server.store.GetUser(ctx, req.Email)
	// if err != nil {
	// 	if err == pgx.ErrNoRows {
	// 		ctx.JSON(http.StatusNotFound, errorResponse(err))
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    req.Email,
	})
}
