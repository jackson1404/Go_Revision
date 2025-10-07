package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	albumModel "jackson.com/goApiDb/internal/albums/models"
	albumService "jackson.com/goApiDb/internal/albums/service"
)

type AlbumHandler struct {
	albumService *albumService.AlbumService
}

func NewAlbumHandler(s *albumService.AlbumService) *AlbumHandler {
	return &AlbumHandler{albumService: s}
}

func (h *AlbumHandler) GetAlbums(c *gin.Context) {

	albums, err := h.albumService.GetAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)

}

func (h *AlbumHandler) AddAlbum(c *gin.Context) {
	var newAlbum albumModel.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.albumService.InsertAlbum(newAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "insert failed: " + err.Error()})
		return
	}
	newAlbum.AlbumID = id
	c.JSON(http.StatusCreated, newAlbum)

}
