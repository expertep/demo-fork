package route

import (
	"demo_fork/controller"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {

	guest := r.Group("/api")

	guest.GET("/transaction/:id", controller.GetTransaction()) // get transaction by id

	guest.POST("/transaction", controller.CreateTransaction())

}
