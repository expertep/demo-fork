package route

import (
	"demo_fork/controller"

	"github.com/gin-gonic/gin"
)

// ROUTE
func Route(r *gin.Engine) {

	quest := r.Group("/api")

	// quest.Use(middlewares.GuestSession())

	quest.GET("/transaction/:id", controller.GetTransaction()) //done

}
