package controllers

import (
	"genshinacademycore/models"
	"genshinacademycore/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type FerretController struct {
	FerretService service.FerretServiceInterface
}

func NewFerretController(serviceFerret service.FerretServiceInterface) FerretController {
	return FerretController{
		FerretService: serviceFerret,
	}
}

func (p FerretController) GetCharactersStats(c *gin.Context) {
	character, err := p.FerretService.GetAllCharactersStats()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": character,
	})
}

func (p FerretController) GetCharacters(c *gin.Context) {
	id := c.Param("id")
	var err error

	var character *[]models.Character

	if id == "" {
		character, err = p.FerretService.GetAllCharacters()
	} else {
		character, err = p.FerretService.GetCharacter(id)
	}

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": character,
	})
}
