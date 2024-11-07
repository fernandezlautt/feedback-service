package feedback

import (
	"fernandezlautt/feedback-service/lib"

	"github.com/gin-gonic/gin"
)

func getAllFeedbacks(c *gin.Context) {
	res, err := findFeedbacks()

	if err != nil {
		lib.AbortWithError(c, err)
		return
	}
	c.JSON(200, res)
}

func FeedbackController(router *gin.RouterGroup) {
	group := router.Group("/feedback")
	group.GET("", getAllFeedbacks)
}
