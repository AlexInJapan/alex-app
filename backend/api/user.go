package api

import (
	"net/http"

	db "github.com/awe8128/backend/db/sqlc"
	"github.com/awe8128/backend/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type createUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	arg := db.CreateUserParams{
		Name:  req.Name,
		Email: req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		dbErr := util.NewDatabaseError(err)
		ErrorResponse(ctx, dbErr)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})

}

type createGetUserRequest struct {
	Email string `form:"email" binding:"required"`
}

func (server *Server) login(ctx *gin.Context) {
	var req createGetUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	user, err := server.store.GetUser(ctx, req.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			dbErr := util.NewDatabaseError(err)
			ErrorResponse(ctx, dbErr)
			return
		}
		ErrorResponse(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
	})
}
