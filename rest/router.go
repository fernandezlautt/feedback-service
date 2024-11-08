package rest

import "fernandezlautt/feedback-service/modules/feedback"

func initRouter() {
	if server == nil {
		panic("Server non existant")
	}
	v1 := server.Group("/v1")
	feedback.FeedbackController(v1)
}
