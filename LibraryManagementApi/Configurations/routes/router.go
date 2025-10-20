package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"jackson.com/libraryapisystem/Configurations/env_configs"
	controllers "jackson.com/libraryapisystem/Controllers"
)

func RegisterRoutes() *http.Server {
	router := gin.Default()

	API_PREFIX := env_configs.AppConfig.API_PREFIX
	API_PORT := env_configs.AppConfig.APP_PORT

	categoryController := controllers.NewCategoryController()

	categoryRouters := router.Group(API_PREFIX + "/categories")
	{
		categoryRouters.POST("", categoryController.CreateCategory)
		categoryRouters.GET("", categoryController.GetAllCategories)
		categoryRouters.GET("/:id", categoryController.GetCategoryByID)
		categoryRouters.PUT("/:id", categoryController.UpdateCategory)
		categoryRouters.DELETE("/:id", categoryController.DeleteCategory)
	}

	authorController := controllers.NewAuthorController()

	authorRoutersGroup := router.Group(API_PREFIX + "/authors")
	{
		authorRoutersGroup.POST("", authorController.CreateAuthor)
		authorRoutersGroup.GET("", authorController.GetAllAuthors)
		authorRoutersGroup.GET("/:id", authorController.GetAuthorByID)
		authorRoutersGroup.PUT("/:id", authorController.UpdateAuthor)
		authorRoutersGroup.DELETE("/:id", authorController.DeleteAuthor)

	}

	bookController := controllers.NewBookController()

	bookRoutersGroup := router.Group(API_PREFIX + "/books")
	{
		bookRoutersGroup.GET("", bookController.GetAllBooks)
		bookRoutersGroup.POST("", bookController.RegisterBook)
		bookRoutersGroup.GET("/:id", bookController.GetBookByID)
		bookRoutersGroup.PUT("/:id", bookController.UpdateBook)
		bookRoutersGroup.DELETE("/:id", bookController.DeleteBook)
	}

	userController := controllers.NewUserController()

	userRoutersGroup := router.Group(API_PREFIX + "/users")
	{
		userRoutersGroup.GET("", userController.GetAllUsers)
		userRoutersGroup.POST("", userController.CreateUser)
		userRoutersGroup.GET("/:id", userController.GetUserByID)
		userRoutersGroup.PUT("/:id", userController.UpdateUser)
		userRoutersGroup.DELETE("/:id", userController.DeleteUser)
	}

	loanController := controllers.NewLoanController()
	reservationController := controllers.NewReservationController()

	loanRoutersGroup := router.Group(API_PREFIX + "/loans")
	{
		loanRoutersGroup.POST("", loanController.LoanBook)
		loanRoutersGroup.PUT("/:id/return", loanController.ReturnBook)

	}

	reservationRoutersGroup := router.Group(API_PREFIX + "/reservations")
	{
		reservationRoutersGroup.POST("", reservationController.ReserveBook)
		reservationRoutersGroup.PUT("/:id/cancel", reservationController.CancelReservation)

	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", API_PORT),
		Handler: router,
	}
	return server
}
