package main

import (
	"github.com/rajakumararul/Agility_Project/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register routes from handlers
	handlers.RegisterRoutes(r)
	handlers.RegisterCountryRoutes(r)
	// Start server on port 8080
	r.Run(":8080")
}
