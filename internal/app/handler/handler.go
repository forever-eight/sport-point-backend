package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/forever-eight/sport-point-backend.git/internal/app/service"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		// registration
		auth.POST("sign-up", h.signUp)
		// authorization
		auth.POST("sign-in", h.signIn)
	}

	/*api := router.Group("/api")
	{
		// classes
		classes := api.Group("/classes")
		{
			// create
			classes.POST("/")
			// get all
			classes.GET("/")
			// get one
			classes.GET("/:id")
			// update
			classes.PUT("/:id")
			//delete
			classes.DELETE("/:id")

			// booking
			booking := classes.Group("/:id/booking")
			{
				booking.POST("/:user-id")
				booking.DELETE("/:id")
			}

		}
		// points
		points := api.Group("/points")
		{
			points.POST("/:user-id")
			points.GET("/:user-id")
		}
	}*/

	return router
}
