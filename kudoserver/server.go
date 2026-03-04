package kudoserver

import (
	"log"
	"main/kudoroutes"

	"github.com/gin-gonic/gin"
)

// StartServer initialises the Gin engine and starts listening.
func StartServer(addr string) error {
	r := gin.Default()

	// Register routes
	kudoroutes.RegisterUserRoutes(r)

	log.Printf("Gin server starting on %s...\n", addr)
	return r.Run(addr)
}
