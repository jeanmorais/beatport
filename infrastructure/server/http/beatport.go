package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) getGenres(c *gin.Context) {

	genres, err := h.beatPortService.GetGenres()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, genres)
}

func (h *handler) getTracksTop10(c *gin.Context) {
	genreKey := c.Param("genreKey")
	genres, err := h.beatPortService.GetTop10(genreKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, genres)
}
