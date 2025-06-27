package handlers

import (
	"github.com/aafeher/gorque/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetSessionList retrieves a list of sessions for a specific user and device from the database and returns them as JSON.
// It requires a valid user ID from the request context and a device ID passed as a query parameter.
// Responds with an error if the user is not found, the device ID is missing, or a database query fails.
func GetSessionList(c *gin.Context) {
	user, ok := GetUserFromContext(c)
	if !ok {
		return
	}

	deviceID := c.Query("device-id")
	if deviceID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "device ID is required"})
		return
	}

	device, err := models.DeviceGetByDeviceID(user.ID, deviceID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "device not found"})
		return
	}

	sessions, err := device.GetSessions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sessions": sessions,
	})
}
