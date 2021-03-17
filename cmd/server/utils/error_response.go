package utils

import (
	"inssa_club_waitlist_backend/cmd/server/errors"

	"github.com/gin-gonic/gin"
)

// AbortWithErrorResponse aborts the request with the given error
func AbortWithErrorResponse(c *gin.Context, statusCode int, errorType string, detail string) {
	c.AbortWithStatusJSON(statusCode, gin.H{"errorType": errorType, "message": errors.Messages[errorType], "detail": detail})
}
