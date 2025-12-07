package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes attaches handler routes to the router
func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", HealthCheck)
	r.GET("/ping", Ping)
}

// HealthCheck returns a 200 OK status with a simple JSON message
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// Ping is a basic GET endpoint returning 200
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
