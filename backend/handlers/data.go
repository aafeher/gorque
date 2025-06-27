package handlers

import (
	"github.com/aafeher/gorque/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetData retrieves time-series data for a specific user, device, and session within a defined time range.
// It verifies the user, extracts query parameters, validates session existence, and queries the InfluxDB instance.
// The response includes data, GPS coordinates, and the center point of the captured coordinates.
func GetData(c *gin.Context) {
	user, ok := GetUserFromContext(c)
	if !ok {
		return
	}

	deviceID := c.Query("device-id")
	if deviceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device ID is required"})
		return
	}

	sessionID := c.Query("session-id")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session not found"})
		return
	}

	session, err := models.SessionGetByIdentifiers(sessionID, deviceID, user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	data, coords, center, err := session.GetSessionData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   data,
		"coords": coords,
		"center": center,
	})
}
