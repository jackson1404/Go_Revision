package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	services "jackson.com/libraryapisystem/Services"
)

type LoanController struct {
	loanService services.LoanService
}

func NewLoanController() *LoanController {
	return &LoanController{
		loanService: services.NewLoanService(),
	}
}

func (ctrl *LoanController) LoanBook(c *gin.Context) {
	var dto services.LoanRequestDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan, err := ctrl.loanService.LoanBook(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, loan)
}

func (ctrl *LoanController) ReturnBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	result, err := ctrl.loanService.ReturnBook(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book returned successfully",
		"data":    result,
	})
}
