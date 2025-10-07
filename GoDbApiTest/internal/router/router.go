package routes

import (
	"github.com/gin-gonic/gin"
	albumHanlder "jackson.com/goApiDb/internal/albums/handlers"
	albumRepo "jackson.com/goApiDb/internal/albums/repository"
	albumService "jackson.com/goApiDb/internal/albums/service"
	config "jackson.com/goApiDb/internal/config"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	repo := albumRepo.NewAlbumRepository(config.DB)
	service := albumService.NewAlbumService(repo)
	handler := albumHanlder.NewAlbumHandler(service)

	albumRoutes := router.Group("/albums")
	{
		albumRoutes.GET("/getAll", handler.GetAlbums)
		albumRoutes.POST("/addAlbum", handler.AddAlbum)
	}

	return router

}
