package assets

import (
	"fmt"
	"ga/internal/types"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type AssetsService interface {
	GetPossibleAssetTypes() []string
	IsValidAssetType(assetType string) bool
	GetPathForAssetType(assetType types.AssetType) string
}

type Controller struct {
	assetsService      AssetsService
	possibleAssetTypes []string
}

func New(assetService AssetsService) *Controller {
	return &Controller{
		assetsService:      assetService,
		possibleAssetTypes: assetService.GetPossibleAssetTypes(),
	}
}

// Upload godoc
//
//	@Summary		Upload assets
//	@Tags			assets
//	@Description	Uploads assets to the specified path.
//	@Accept			multipart/form-data
//	@Param			path	formData	string	true	"Path to upload files" Enums(characters, characters/icons, tables, news, opengraph))
//	@Param			files	formData	[]file	true	"Files to upload"
//	@Security		ApiKeyAuth
//	@Router			/assets/{path} [post]
//	@Success		200
//	@Failure		400
func (c *Controller) Upload(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "files not specified"})
		return
	}

	assetType := strings.Trim(ctx.Param("path"), "/")
	if !c.assetsService.IsValidAssetType(assetType) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":    "wrong file type",
			"possible": c.possibleAssetTypes,
		})
		return
	}

	var (
		successful []string
		errors     []string
	)

	assetPath := c.assetsService.GetPathForAssetType(types.AssetType(assetType))
	for _, file := range files {
		savePath := filepath.Join(assetPath, file.Filename)
		if err := ctx.SaveUploadedFile(file, savePath); err != nil {
			errors = append(errors, fmt.Sprintf("failed to upload %s: %s", file.Filename, err.Error()))
		} else {
			successful = append(successful, file.Filename)
		}
	}

	ctx.JSON(assetsResult(successful, errors))
}

// Delete godoc
//
//	@Summary		Delete assets
//	@Tags			assets
//	@Description	Deletes assets at the specified paths.
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			paths	formData	[]string	true	"Assets paths to delete"	collectionFormat(multi)	example(characters/icons/lisa.webp,characters/icons/hu_tao.webp)
//	@Security		ApiKeyAuth
//	@Success		200
//	@Failure		400
//	@Router			/assets [delete]
func (c *Controller) Delete(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paths := form.Value["paths"]
	if len(paths) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "paths not specified"})
		return
	}

	var (
		successful []string
		errors     []string
	)

	for _, path := range paths {
		parts := strings.Split(path, "/")
		assetType := strings.Join(parts[:len(parts)-1], "/")
		fileName := parts[len(parts)-1]

		if !c.assetsService.IsValidAssetType(assetType) {
			errors = append(errors, fmt.Sprintf("wrong asset type: %s", assetType))
			continue
		}

		if err := os.Remove(filepath.Join(c.assetsService.GetPathForAssetType(types.AssetType(assetType)), fileName)); err != nil {
			errors = append(errors, fmt.Sprintf("failed to delete %s: %s", path, err.Error()))
			continue
		}

		successful = append(successful, path)
	}

	ctx.JSON(assetsResult(successful, errors))
}

func assetsResult(successful []string, errors []string) (status int, response gin.H) {
	if len(successful) > 0 && len(errors) > 0 {
		response = gin.H{
			"successful": successful,
			"errors":     errors,
		}
		status = http.StatusOK
	} else if len(errors) > 0 {
		response = gin.H{
			"errors": errors,
		}
		status = http.StatusInternalServerError
	} else {
		response = gin.H{
			"successful": successful,
		}
		status = http.StatusOK
	}

	return
}
