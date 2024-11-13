package lib

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func AbortWithError(c *gin.Context, err error) {
	c.Error(err)
	c.Abort()
}

func GetCorrelationId(c *gin.Context) string {
	value := c.GetHeader("correlation_id")

	if len(value) == 0 {
		value = uuid.NewV4().String()
	}

	return value
}

type Pagination struct {
	Skip  int64
	Limit int64
}

func GetPagination(page int, size int) *Pagination {
	if page < 1 {
		page = 1
	}

	if size < 1 {
		size = 10
	}

	return &Pagination{
		Skip:  int64((page - 1) * size),
		Limit: int64(size),
	}
}
