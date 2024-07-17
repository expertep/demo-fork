package route

import (
	"demo_fork/controller"
	middleware "demo_fork/middlewares"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {

	guest := r.Group("/api")

	guest.GET("/transaction/:id", middleware.Auth(), controller.GetTransaction()) // get transaction by id

	guest.POST("/transaction", controller.CreateTransaction())

}
