package daos

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"jackson.com/libraryapisystem/Configurations/db_configs"
	"jackson.com/libraryapisystem/models"
)

type ReservationDAO interface {
	CreateReservation(res *models.Reservation, tx *gorm.DB) error
	CancelReservation(id uint) error
	GetDB() *gorm.DB
}

type ReservationDAOImpl struct {
	db *gorm.DB
}

func NewReservationDAO() *ReservationDAOImpl {
	return &ReservationDAOImpl{db: db_configs.DB}
}

func (dao *ReservationDAOImpl) CreateReservation(res *models.Reservation, tx *gorm.DB) error {
	return tx.Create(res).Error
}

func (dao *ReservationDAOImpl) GetDB() *gorm.DB {
	return dao.db
}

func (dao *ReservationDAOImpl) CancelReservation(id uint) error {
	var res models.Reservation
	if err := dao.db.First(&res, id).Error; err != nil {
		return errors.New("reservation not found")
	}

	now := time.Now()
	res.ExpiresAt = &now
	return dao.db.Save(&res).Error
}
