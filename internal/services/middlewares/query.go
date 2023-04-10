package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLimitOffset() gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil || limit < 1 {
			limit = 1000
		}

		offset, err := strconv.Atoi(c.Query("offset"))
		if err != nil || offset < 0 {
			offset = 0
		}

		c.Set("limit", uint(limit))
		c.Set("offset", uint(offset))
		c.Next()
	}
}
