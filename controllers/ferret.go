package controllers

import (
	"genshinacademycore/models"
	models "genshinacademycore/models/db"
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

func (с FerretController) GetCharactersStats(c *gin.Context) {
	character, err := с.FerretService.GetAllCharactersStats()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": character,
	})
}

func (с FerretController) GetCharacters(c *gin.Context) {
	id := c.Param("id")
	var err error

	var character *[]models.Character

	if id == "" {
		character, err = с.FerretService.GetAllCharacters()
	} else {
		character, err = с.FerretService.GetCharacter(id)
	}

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": character,
	})
}
