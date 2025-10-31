package handler

import (
	"github.com/NQFV/p/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		transaction := api.Group("/transaction")
		{
			transaction.POST("/", h.createTransaction)
			transaction.GET("/", h.getAllTransaction)
			transaction.GET("/:id", h.getByIdTransaction)
			transaction.PUT("/:id", h.updateTransaction)
			transaction.DELETE("/:id", h.deleteTransaction)
		}

		category := api.Group("/category")
		{
			category.POST("/", h.createCategory)
			category.GET("/", h.getCategory)
			category.PUT("/:category_id", h.updateCategory)
			category.DELETE("/:category_id", h.deleteCategory)
		}

		anal := api.Group("/anal")
		{
			anal.GET("/:id", h.anal)
		}
	}

	return router
}
