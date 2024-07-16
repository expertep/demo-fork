package route

import (
	"demo_fork/controller"

	"github.com/gin-gonic/gin"
)

// ROUTE
func Route(r *gin.Engine) {

	guest := r.Group("/api")

	guest.GET("/transaction/:id", controller.GetTransaction())

}
