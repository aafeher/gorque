package handlers

import (
	"github.com/aafeher/gorque/middlewares"
	"github.com/aafeher/gorque/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// Register is a handler function that registers a new user with email and hashed password and saves it in the database.
func Register(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	if err := validateEmail(body.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validatePassword(body.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password hash generation failed"})
		return
	}
	user := models.User{Email: body.Email, Password: string(hashed)}
	if err = models.UserCreate(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// Login handles user authentication by verifying email and password, generating a JWT token, and responding with the token.
func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	user, err := models.UserGetByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	tokenStr, err := middlewares.GenerateJWT(user.ID, 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}

// GetProfile retrieves the profile information of the currently authenticated user from the database and returns it in JSON format.
func GetProfile(c *gin.Context) {
	user, ok := GetUserFromContext(c)
	if !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": user.Email, "id": user.ID, "name": user.Name})
}

// UpdateProfileName updates the name of a user profile identified by the ID in the request URL.
// It expects a JSON body containing the new name and returns appropriate HTTP responses based on the operation outcome.
func UpdateProfileName(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	if err := validateName(body.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, ok := GetUserFromContext(c)
	if !ok {
		return
	}

	if err := user.UpdateName(body.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Name updated successfully"})
}

// UpdateProfilePassword updates the password of a user based on the provided current and new password.
// It validates the current password, hashes the new password, and updates it in the database.
func UpdateProfilePassword(c *gin.Context) {
	var body struct {
		CurrentPassword string `json:"currentPassword"`
		NewPassword     string `json:"newPassword"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	if err := validatePassword(body.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "new " + err.Error()})
		return
	}

	if body.CurrentPassword == body.NewPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "new password must be different from current password"})
		return
	}

	user, ok := GetUserFromContext(c)
	if !ok {
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.CurrentPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "current password is incorrect"})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password hash generation failed"})
		return
	}

	if err := user.UpdatePassword(string(hashed)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

// RefreshToken generates a new JWT token for the authenticated user if the current token is valid.
// Returns an error response if the userID is invalid or token generation fails.
func RefreshToken(c *gin.Context) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	userID, ok := userIDRaw.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}

	token, err := middlewares.GenerateJWT(userID, 1*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
