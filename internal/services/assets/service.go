package assets

import (
	"fmt"
	"ga/internal/academy_core"
	"net/http"
	"os"
	"path/filepath"
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
	Characters      AssetsType = "characters"
	CharactersIcons AssetsType = Characters + "/icons"
	Tables          AssetsType = "tables"
	News            AssetsType = "news"
	OpenGraph       AssetsType = "opengraph"
)

var validAssetTypes = []AssetsType{
	Characters,
	CharactersIcons,
	Tables,
	News,
	OpenGraph,
}

// UploadAssets godoc
//
//	@Summary		Upload assets
//	@Tags			assets
//	@Description	Uploads assets to the specified path.
//	@Description	Possible values:
//	@Description	* characters
//	@Description	* characters/icons
//	@Description	* tables
//	@Description	* news
//	@Description	* opengraph
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			path	path		string	true	"Path to upload files. Possible values: characters, characters/icons, tables, news, opengraph"
//	@Param			files	formData	file	true	"Files to upload"
//	@Security		ApiKeyAuth
//	@Router			/assets/{path} [post]
//	@Success		200	{object}	gin.H{}
//	@Failure		400	{object}	gin.H{}
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
	if !isValidAssetType(path) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "wrong file type",
			"possible": validAssetTypes,
		})
		return
	}

	var (
		successful []string
		errors     []string
	)

	for _, file := range files {
		savePath := filepath.Join(service.assetsPath, path, file.Filename)
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			errors = append(errors, fmt.Sprintf("failed to upload %s: %s", file.Filename, err.Error()))
		} else {
			successful = append(successful, file.Filename)
		}
	}

	assetsResult(c, successful, errors)
}

// DeleteAssets godoc
//
//	@Summary		Delete assets
//	@Tags			assets
//	@Description	Deletes assets at the specified paths.
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			paths	formData	[]string	true	"Assets paths to delete"	collectionFormat(multi)	example(characters/icons/lisa.webp)
//	@Security		ApiKeyAuth
//	@Router			/assets [delete]
//	@Success		200	{object}	gin.H{}
//	@Failure		400	{object}	gin.H{}
func (service *Service) Delete(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paths := form.Value["paths"]
	if len(paths) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "paths not specified"})
		return
	}

	var (
		successful []string
		errors     []string
	)

	for _, path := range paths {
		if err := os.Remove(filepath.Join(service.assetsPath, path)); err != nil {
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

func isValidAssetType(assetType string) bool {
	for _, validType := range validAssetTypes {
		if assetType == string(validType) {
			return true
		}
	}
	return false
}
