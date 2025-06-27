package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetDeviceList retrieves the list of devices associated with the authenticated user and returns them in the response.
func GetDeviceList(c *gin.Context) {
	user, ok := GetUserFromContext(c)
	if !ok {
		return
	}

	devices, err := user.GetDevices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"devices": devices,
	})
}
