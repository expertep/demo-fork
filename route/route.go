package route

import (
	"demo_fork/controller"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {

	guest := r.Group("/api")

	guest.GET("/transaction/:id", controller.GetTransaction())

	quest.POST("/transaction", controller.CreateTransaction()) //done

}
