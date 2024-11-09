package feedback

import (
	"fernandezlautt/feedback-service/lib"
	"fernandezlautt/feedback-service/security"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAllFeedbacksController(c *gin.Context) {

	articleId := c.Param("articleId")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	res, err := findFeedbacks(articleId, page, size)

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

	articleId := c.Param("articleId")
	// get user and take the name
	user_rec, _ := c.Get("user")
	user, _ := user_rec.(security.User)

	err := createFeedback(c, feedbackDto, user.Name, articleId)

	if err != nil {
		lib.AbortWithError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Feedback created",
	})
}

func FeedbackController(router *gin.RouterGroup) {
	group := router.Group("/feedback/:articleId")
	group.GET("", getAllFeedbacksController)
	group.POST("", createFeedbackController)
}
