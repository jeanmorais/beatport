package http

import (
	"github.com/jeanmorais/beatport/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	beatPortService domain.BeatPortService
}

// NewHandler creates a new handler
func NewHandler(beatPortService domain.BeatPortService) http.Handler {
	handler := &handler{
		beatPortService: beatPortService,
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger(), handler.recovery())

	router.GET("/tracks/top10/:genreKey", handler.getTracksTop10)
	router.GET("/genres", handler.getGenres)

	return router
}

func (h *handler) recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
