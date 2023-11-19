package handlers

import (
	"github.com/gin-gonic/gin"
)

func BuildError(err string, msg any) gin.H {
	var e = make(gin.H)
	e["error"] = err
	if msg != nil {
		e["message"] = msg
	}
	return e
}
