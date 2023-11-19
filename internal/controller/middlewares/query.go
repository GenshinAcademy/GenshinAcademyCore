package middlewares

import (
	"github.com/gin-gonic/gin"
	"strconv"
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

		c.Set("limit", limit)
		c.Set("offset", offset)
		c.Next()
	}
}

func GetSort() gin.HandlerFunc {
	return func(c *gin.Context) {
		sort := c.Query("sort")

		c.Set("sort", sort)
		c.Next()
	}
}
