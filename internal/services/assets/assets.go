package assets

import (
	"fmt"
	"ga/internal/academy_core"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Service struct {
	core       *academy_core.AcademyCore
	assetsPath string
}

func CreateService(core *academy_core.AcademyCore, assetsPath string) *Service {
	var result *Service = new(Service)
	result.core = core
	result.assetsPath = assetsPath
	return result
}

type AssetsType string

const (
	characters      AssetsType = "characters"
	charactersIcons AssetsType = characters + "/icons"
	tables          AssetsType = "tables"
	news            AssetsType = "news"
)

var possibleTypes []AssetsType = []AssetsType{characters, charactersIcons, tables, news}

func (service *Service) Upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "files not specified"})
		return
	}

	path := strings.Trim(c.Param("path"), "/")
	if !isValidAssetType(AssetsType(path)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "wrong file type",
			"possible": possibleTypes,
		})
		return
	}

	for _, file := range files {
		savePath := fmt.Sprintf("%s/%s/%s", service.assetsPath, path, file.Filename)
		fmt.Println(savePath)
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, files)
}

func (service *Service) Delete(c *gin.Context) {
	path := strings.Trim(c.Param("path"), "/")

	if err := os.Remove(fmt.Sprintf("%s/%s", service.assetsPath, path)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func isValidAssetType(assetType AssetsType) bool {
	switch assetType {
	case characters, charactersIcons, tables, news:
		return true
	default:
		return false
	}
}
