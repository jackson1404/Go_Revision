package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	services "jackson.com/libraryapisystem/Services"
)

type ReservationController struct {
	reservationService services.ReservationService
}

func NewReservationController() *ReservationController {
	return &ReservationController{
		reservationService: services.NewReservationService(),
	}
}

func (ctrl *ReservationController) ReserveBook(c *gin.Context) {
	var dto services.ReservationRequestDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ctrl.reservationService.ReserveBook(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (ctrl *ReservationController) CancelReservation(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid reservation id"})
		return
	}

	if err := ctrl.reservationService.CancelReservation(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reservation cancelled"})
}
