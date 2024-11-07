package rest

import (
	"fernandezlautt/feedback-service/lib"
	"fmt"
)

func Init() {
	createServer()
	initRouter()
	getServer().Run(fmt.Sprintf(":%d", lib.GetEnv().Port))
}
