package route

import (
	"demo_fork/controller"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {

	guest := r.Group("/api")

	guest.GET("/transaction/:id", controller.GetTransaction()) // get transaction by id
	guest.GET("/transaction/:id/1", controller.GetTransaction())
	guest.GET("/transaction/:id/3", controller.GetTransaction())
	guest.GET("/transaction/:id/2", controller.GetTransaction())
	guest.GET("/transaction/:id/4", controller.GetTransaction())

	guest.POST("/transaction", controller.CreateTransaction())

}
