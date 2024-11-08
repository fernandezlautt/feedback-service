package feedback

import (
	"fernandezlautt/feedback-service/lib"

	"github.com/gin-gonic/gin"
)

func getAllFeedbacksController(c *gin.Context) {
	res, err := findFeedbacks()

	if err != nil {
		lib.AbortWithError(c, err)
		return
	}
	c.JSON(200, res)
}

func createFeedbackController(c *gin.Context) {
	var feedbackDto CreateFeedbackDto
	if err := c.ShouldBindJSON(&feedbackDto); err != nil {
		lib.AbortWithError(c, err)
		return
	}

	err := createFeedback(c, feedbackDto)
	if err != nil {
		lib.AbortWithError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Feedback created",
	})
}

func FeedbackController(router *gin.RouterGroup) {
	group := router.Group("/feedback")
	group.GET("", getAllFeedbacksController)
	group.POST("", createFeedbackController)
}
