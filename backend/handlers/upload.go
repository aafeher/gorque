package handlers

import (
	"github.com/gin-gonic/gin"
)

// uploadService is the global instance of the upload service
var uploadService = NewUploadService()

// Upload handles file upload requests by delegating to the upload service
// This function maintains backward compatibility while using the new refactored service architecture
func Upload(c *gin.Context) {
	uploadService.ProcessUpload(c)
}
