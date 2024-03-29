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

	var (
		successful []string
		errors     []string
	)

	for _, file := range files {
		savePath := fmt.Sprintf("%s/%s/%s", service.assetsPath, path, file.Filename)
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			errors = append(errors, fmt.Sprintf("failed to upload %s: %s", file.Filename, err.Error()))
		} else {
			successful = append(successful, file.Filename)
		}
	}

	assetsResult(c, successful, errors)
}

func (service *Service) Delete(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paths := form.Value["path"]
	if len(paths) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paths not specified"})
		return
	}

	var (
		successful []string
		errors     []string
	)

	for _, path := range paths {
		if err := os.Remove(fmt.Sprintf("%s/%s", service.assetsPath, path)); err != nil {
			errors = append(errors, fmt.Sprintf("failed to delete %s: %s", path, err.Error()))
		} else {
			successful = append(successful, path)
		}
	}

	assetsResult(c, successful, errors)
}

func assetsResult(c *gin.Context, successful []string, errors []string) {
	if len(successful) > 0 && len(errors) > 0 {
		response := gin.H{
			"successful": successful,
			"errors":     errors,
		}
		c.JSON(http.StatusOK, response)
	} else if len(errors) > 0 {
		response := gin.H{
			"errors": errors,
		}
		c.JSON(http.StatusInternalServerError, response)
	} else {
		response := gin.H{
			"successful": successful,
		}
		c.JSON(http.StatusOK, response)
	}
}

func isValidAssetType(assetType AssetsType) bool {
	switch assetType {
	case characters, charactersIcons, tables, news:
		return true
	default:
		return false
	}
}
