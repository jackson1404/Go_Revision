package daos

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"jackson.com/libraryapisystem/Configurations/db_configs"
	"jackson.com/libraryapisystem/models"
)

type LoanDAO interface {
	CreateLoan(loan *models.Loan, tx *gorm.DB) error
	GetLoanByID(id uint) (*models.Loan, error)
	FindLoanByID(id uint) (*models.Loan, error)
	MarkLoanAsReturned(loan *models.Loan, fineCents int64) error
	GetDB() *gorm.DB
}

type LoanDAOImpl struct {
	db *gorm.DB
}

func NewLoanDAO() *LoanDAOImpl {
	return &LoanDAOImpl{db: db_configs.DB}
}

func (dao *LoanDAOImpl) CreateLoan(loan *models.Loan, tx *gorm.DB) error {
	return tx.Create(loan).Error
}

func (dao *LoanDAOImpl) GetLoanByID(id uint) (*models.Loan, error) {
	var loan models.Loan
	err := dao.db.First(&loan, id).Error
	return &loan, err
}

func (dao *LoanDAOImpl) GetDB() *gorm.DB {
	return dao.db
}

func (dao *LoanDAOImpl) FindLoanByID(id uint) (*models.Loan, error) {
	var loan models.Loan
	err := dao.db.First(&loan, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("loan not found")
	}
	return &loan, err
}

func (dao *LoanDAOImpl) MarkLoanAsReturned(loan *models.Loan, fineCents int64) error {
	now := time.Now()
	loan.ReturnedAt = &now
	loan.FineCents = fineCents
	return dao.db.Save(loan).Error
}
