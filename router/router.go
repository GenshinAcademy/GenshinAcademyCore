package router

import (
	"genshinacademycore/controllers"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterController struct {
	Ferret controllers.FerretController
}

func NewRouter(controller RouterController) *gin.Engine {
	r := gin.New()

	mode, _ := os.LookupEnv("ENV")
	gin.SetMode(mode)

	r.Use(cors.Default())

	character := r.Group("/characters")
	{
		// character.GET("characters/", controller.Ferret.GetCharacters)
		// character.GET("characters/:id", controller.Ferret.GetCharacters)
		character.GET("stats", controller.Ferret.GetCharactersStats)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found!",
		})
	})

	return r
}
