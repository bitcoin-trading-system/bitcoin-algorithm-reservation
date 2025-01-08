package router

import (
	"net/http"

	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/config"
	"github.com/bitcoin-trading-system/bitcoin-algorithm-reservation/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg config.Config) *gin.Engine {
	h := handler.NewHandler(cfg)

	r := gin.Default()

	return setUpRouter(r, h)
}

func setUpRouter(r *gin.Engine, h handler.IHandler) *gin.Engine {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.POST("/reservations", h.PostReservation)
	r.GET("/reservations/:id", h.GetReservationByID)
	r.GET("/reservations", h.GetReservations)
	r.PUT("/reservations/:id", h.UpdateReservation)
	r.DELETE("/reservations/:id", h.DeleteReservation)

	return r
}
