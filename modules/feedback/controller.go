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

	// service
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
	user_rec, _ := c.Get("user")

	// type assertion
	user, _ := user_rec.(security.User)
	// service
	err := createFeedback(c, feedbackDto, &user, articleId)

	if err != nil {
		lib.AbortWithError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Feedback created",
	})
}

func disableFeedbackController(c *gin.Context) {
	feedbackId := c.Param("feedbackId")

	var disableFeedbackDto DisableFeedbackDto

	if err := c.ShouldBindJSON(&disableFeedbackDto); err != nil {
		lib.AbortWithError(c, err)
		return
	}

	// service
	err := disableFeedback(feedbackId, disableFeedbackDto.Reason)

	if err != nil {
		lib.AbortWithError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Feedback disabled",
	})
}

// as the router is a reference to the main gin.RouterGroup, we can modify it directly witoout returning anything
func FeedbackController(router *gin.RouterGroup) {
	group := router.Group("/feedback")
	group.GET("/:articleId", getAllFeedbacksController)
	group.POST("/:articleId", createFeedbackController)
	group.POST("/disable/:feedbackId", disableFeedbackController)
}
