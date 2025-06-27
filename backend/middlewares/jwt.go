package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"time"
)

// JWTKey holds the secret key used for signing and verifying JWT tokens, retrieved from the environment variable.
var JWTKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// JWTAuthMiddleware is a middleware that validates the Authorization header, parses the JWT token, and verifies its claims.
// If the token is valid, it extracts the user_id claim and stores it in the context for further processing.
// Invalid or missing tokens result in appropriate HTTP 401 or 500 error responses and abort the request.
func JWTAuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong Authorization method"})
		return
	}

	tokenStr := parts[1]
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return JWTKey, nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["user_id"] == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing user_id claim"})
		return
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Wrong user_id claim type"})
		return
	}
	c.Set("userID", uint(userIDFloat))
	c.Next()
}

// GenerateJWT creates and returns a signed JSON Web Token (JWT) containing the specified user ID as a claim.
// The token includes an expiration time defined by the expirationTime parameter.
// Returns the signed token string or an error if the signing process fails.
func GenerateJWT(userID uint, expirationTime time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expirationTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTKey)
}
