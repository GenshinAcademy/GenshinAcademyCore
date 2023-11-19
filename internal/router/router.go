package router

import (
	"fmt"
	"ga/internal/controller/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CharactersController interface {
	GetAll(c *gin.Context)
	Create(c *gin.Context)
}

type WeaselAppraiserController interface {
	GetAll(c *gin.Context)
	UpdateStats(c *gin.Context)
}

type NewsController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetAll(c *gin.Context)
}

type TablesController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetAll(c *gin.Context)
}

type AssetsController interface {
	Upload(c *gin.Context)
	Delete(c *gin.Context)
}

type Router struct {
	Engine    *gin.Engine
	authToken string

	apiGroup  *gin.RouterGroup
	apiGroups map[string]*gin.RouterGroup
}

func New(
	ginMode string,
	authToken string,
) *Router {
	r := &Router{
		authToken: authToken,
	}

	gin.SetMode(ginMode)
	r.Engine = gin.Default()

	r.Engine.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Accept-Languages"},
		MaxAge:          12 * time.Hour,
	}))

	r.apiGroups = make(map[string]*gin.RouterGroup)
	r.apiGroup = r.Engine.Group("/api")

	r.Engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
	})

	return r
}

func (r *Router) Run(port uint16) error {
	return r.Engine.Run(fmt.Sprintf(":%d", port))
}

func (r *Router) WithCharactersController(charactersController CharactersController) *Router {
	characters := r.getGroup("/characters")
	{
		characters.GET("/", middlewares.GetLimitOffset(), charactersController.GetAll)
		characters.POST("/", middlewares.Authenticate(r.authToken), charactersController.Create)
	}
	return r
}

func (r *Router) WithWeaselAppraiserController(weaselAppraiserController WeaselAppraiserController) *Router {
	characters := r.getGroup("/characters")
	{
		characters.GET("/stats", middlewares.GetLimitOffset(), weaselAppraiserController.GetAll)
		characters.PATCH("/stats/:id", middlewares.Authenticate(r.authToken), weaselAppraiserController.UpdateStats)
	}
	return r
}

func (r *Router) WithNewsController(newsController NewsController) *Router {
	newsGroup := r.getGroup("/news")
	{
		newsGroup.GET("/", middlewares.GetLimitOffset(), middlewares.GetSort(), newsController.GetAll)
		newsGroup.POST("/", middlewares.Authenticate(r.authToken), newsController.Create)
		newsGroup.PATCH("/:id", middlewares.Authenticate(r.authToken), newsController.Update)
	}
	return r
}

func (r *Router) WithTablesController(tablesController TablesController) *Router {
	tables := r.getGroup("/tables")
	{
		tables.GET("/", middlewares.GetLimitOffset(), middlewares.GetSort(), tablesController.GetAll)
		tables.POST("/", middlewares.Authenticate(r.authToken), tablesController.Create)
		tables.PATCH("/:id", middlewares.Authenticate(r.authToken), tablesController.Update)
	}
	return r
}

func (r *Router) WithAssetsController(assetsController AssetsController) *Router {
	assets := r.getGroup("/assets")
	{
		assets.POST("/*path", middlewares.Authenticate(r.authToken), assetsController.Upload)
		assets.DELETE("/*path", middlewares.Authenticate(r.authToken), assetsController.Delete)
	}
	return r
}

func (r *Router) getGroup(path string) *gin.RouterGroup {
	if group, ok := r.apiGroups[path]; ok {
		return group
	}

	group := r.apiGroup.Group(path)
	r.apiGroups[path] = group

	return group
}
