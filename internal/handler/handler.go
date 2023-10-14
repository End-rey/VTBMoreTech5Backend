package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "github.com/End-rey/VTBMoreTech5Backend/docs"
	"github.com/End-rey/VTBMoreTech5Backend/internal/service"
	"github.com/End-rey/VTBMoreTech5Backend/pkg/logger"
)

type Handler struct {
	services *service.Services
	l logger.Interface
}

func NewHandler(services *service.Services, l logger.Interface) *Handler {
	return &Handler{services: services, l:l}
}

// NewRouter -.
// Swagger spec:
// @title       VTB MoreTech 5.0
// @description The problem of finding an optimal branch of VTB Bank
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func (h *Handler) NewRouter() *gin.Engine {
	router := gin.New()
	// Options
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// K8s probe
	router.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Routers
	api := router.Group("/api")
	{
		office := api.Group("/offices")
		{
			office.GET("/", h.getAllOffices)
		}

		atm := api.Group("/atms")
		{
			atm.GET("/", h.getAllAtms)
		}
	}

	return router
}
