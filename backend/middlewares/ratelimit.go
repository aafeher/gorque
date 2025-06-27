package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

// ClientInfo holds rate limiting information for a client
type ClientInfo struct {
	requests  int       // number of requests
	lastReset time.Time // last time the counter was reset
	blocked   time.Time // when the client was blocked
}

// RateLimiter holds the rate limiting configuration and client data
type RateLimiter struct {
	clients   map[string]*ClientInfo
	mutex     sync.RWMutex
	maxReqs   int           // maximum requests per window
	window    time.Duration // time window
	blockTime time.Duration // how long to block after exceeding limit
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(maxReqs int, window time.Duration, blockTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		clients:   make(map[string]*ClientInfo),
		maxReqs:   maxReqs,
		window:    window,
		blockTime: blockTime,
	}

	// Cleanup goroutine to remove old entries
	go rl.cleanup()

	return rl
}

// cleanup removes old client entries periodically
func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rl.mutex.Lock()
			now := time.Now()
			for ip, client := range rl.clients {
				if now.Sub(client.lastReset) > rl.window*2 {
					delete(rl.clients, ip)
				}
			}
			rl.mutex.Unlock()
		}
	}
}

// isAllowed checks if a request from the given IP is allowed
func (rl *RateLimiter) isAllowed(ip string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	client, exists := rl.clients[ip]

	if !exists {
		rl.clients[ip] = &ClientInfo{
			requests:  1,
			lastReset: now,
		}
		return true
	}

	// Check if client is currently blocked
	if !client.blocked.IsZero() && now.Sub(client.blocked) < rl.blockTime {
		return false
	}

	// Reset window if needed
	if now.Sub(client.lastReset) >= rl.window {
		client.requests = 1
		client.lastReset = now
		client.blocked = time.Time{} // Clear block
		return true
	}

	// Increment requests
	client.requests++

	// Check if limit exceeded
	if client.requests > rl.maxReqs {
		client.blocked = now
		return false
	}

	return true
}

// RateLimitMiddleware creates a middleware function for rate limiting
func RateLimitMiddleware(maxReqs int, window time.Duration, blockTime time.Duration) gin.HandlerFunc {
	limiter := NewRateLimiter(maxReqs, window, blockTime)

	return func(c *gin.Context) {
		ip := c.ClientIP()

		if !limiter.isAllowed(ip) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests, please try again later",
			})
			return
		}

		c.Next()
	}
}
