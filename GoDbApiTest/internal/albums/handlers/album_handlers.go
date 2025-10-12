package handlers

import (
	"net/http"
	"strconv"

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

// PUT /albums/update
func (h *AlbumHandler) UpdateAlbum(c *gin.Context) {
	var album albumModel.Album
	if err := c.BindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.albumService.UpdateAlbum(album); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "album updated successfully"})
}

func (h *AlbumHandler) DeleteAlbum(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid album ID"})
		return
	}

	if err := h.albumService.DeleteAlbum(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "album deleted successfully"})
}
