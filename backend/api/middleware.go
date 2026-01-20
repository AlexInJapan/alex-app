package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/awe8128/backend/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func ErrorHandler() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			log.Printf("Panic recovered: %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}

func ErrorResponse(c *gin.Context, err error) {
	var appErr *util.AppError

	if errors.As(err, &appErr) {
		c.JSON(appErr.StatusCode, gin.H{
			"error": gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
				"details": appErr.Details,
			},
		})
		return
	}

	if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": gin.H{
				"code":    40401,
				"message": "Resource not found",
			},
		})
		return
	}
	if c.IsAborted() {
		return
	}
	// Default internal server errro
	log.Printf("Unhandled error: %v", err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": gin.H{
			"code":    50002,
			"message": "Internal server error",
		},
	})
}
