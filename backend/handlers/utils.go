package handlers

import (
	"github.com/aafeher/gorque/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserFromContext extracts the user ID from the context and retrieves the user object.
// Returns the user object and true if successful, or nil and false if failed.
// It automatically sends an appropriate error response to the client in case of failure.
func GetUserFromContext(c *gin.Context) (*models.User, bool) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user ID not found in context"})
		return nil, false
	}

	id, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID type"})
		return nil, false
	}

	user, err := models.UserGetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return nil, false
	}

	return user, true
}
