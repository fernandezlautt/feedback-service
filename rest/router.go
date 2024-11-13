package rest

import "fernandezlautt/feedback-service/modules/feedback"

func initRouter() {
	if server == nil {
		panic("Server non existant")
	}
	// proptected middleware check bearer, btw checking bearer instead of Bearer is a little bit annoying
	v1 := server.Group("/v1", ProtectedMiddleware)
	feedback.FeedbackController(v1)
}
