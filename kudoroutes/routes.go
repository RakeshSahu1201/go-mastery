package kudoroutes

import (
	"main/kudohandlers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registers all /users endpoints.
func RegisterUserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", kudohandlers.GetUsers)
		users.POST("/", kudohandlers.CreateUser)
		users.GET("/:id", kudohandlers.GetUserByID)
	}
}
