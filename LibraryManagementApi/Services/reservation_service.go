package services

import (
	"errors"
	"time"

	"gorm.io/gorm"
	daos "jackson.com/libraryapisystem/DAOs"
	"jackson.com/libraryapisystem/models"
)

type ReservationRequestDTO struct {
	UserID    uint       `json:"user_id" binding:"required"`
	BookID    uint       `json:"book_id" binding:"required"`
	ExpiresAt *time.Time `json:"expires_at"`
}

type ReservationResponseDTO struct {
	ID         uint       `json:"id"`
	UserID     uint       `json:"user_id"`
	BookID     uint       `json:"book_id"`
	ReservedAt time.Time  `json:"reserved_at"`
	Notified   bool       `json:"notified"`
	ExpiresAt  *time.Time `json:"expires_at"`
}

type ReservationService interface {
	ReserveBook(dto *ReservationRequestDTO) (*ReservationResponseDTO, error)
	CancelReservation(reservationID uint) error
}

type ReservationServiceImpl struct {
	reservationDAO daos.ReservationDAO
	bookDAO        daos.BookDAO
	userDAO        daos.UserDAO
}

func NewReservationService() *ReservationServiceImpl {
	return &ReservationServiceImpl{
		reservationDAO: daos.NewReservationDAO(),
		bookDAO:        daos.NewBookDAO(),
		userDAO:        daos.NewUserDAO(),
	}
}

func (s *ReservationServiceImpl) ReserveBook(dto *ReservationRequestDTO) (*ReservationResponseDTO, error) {
	returnValue := &ReservationResponseDTO{}
	err := s.bookDAO.GetDB().Transaction(func(tx *gorm.DB) error {

		user, err := s.userDAO.FindUserByID(dto.UserID)
		if err != nil {
			return errors.New("user not found")
		}

		book, err := s.bookDAO.GetBookByID(dto.BookID)
		if err != nil {
			return errors.New("book not found")
		}

		res := &models.Reservation{
			UserID:     user.ID,
			BookID:     book.ID,
			ReservedAt: time.Now(),
			ExpiresAt:  dto.ExpiresAt,
			Notified:   false,
		}
		if err := s.reservationDAO.CreateReservation(res, tx); err != nil {
			return err
		}

		*returnValue = ReservationResponseDTO{
			ID:         res.ID,
			UserID:     res.UserID,
			BookID:     res.BookID,
			ReservedAt: res.ReservedAt,
			ExpiresAt:  res.ExpiresAt,
			Notified:   res.Notified,
		}
		return nil
	})
	return returnValue, err
}

func (s *ReservationServiceImpl) CancelReservation(reservationID uint) error {
	return s.reservationDAO.CancelReservation(reservationID)
}
