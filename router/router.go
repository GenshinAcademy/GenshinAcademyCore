package router

import (
	"genshinacademycore/controllers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type RouterController struct {
	Ferret controllers.FerretController
}

func NewRouter(controller RouterController) *gin.Engine {
	r := gin.New()

	mode, _ := os.LookupEnv("GIN_MODE")
	gin.SetMode(mode)

	character := r.Group("/api/v1/")
	{
		character.GET("characters/", controller.Ferret.GetCharacters)
		character.GET("characters/:id", controller.Ferret.GetCharacters)
		character.GET("characters/stats", controller.Ferret.GetCharactersStats)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Page not found!",
		})
	})

	return r
}
