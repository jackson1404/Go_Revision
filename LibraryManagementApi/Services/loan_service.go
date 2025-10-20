package services

import (
	"errors"
	"time"

	"gorm.io/gorm"
	daos "jackson.com/libraryapisystem/DAOs"
	"jackson.com/libraryapisystem/models"
)

type LoanRequestDTO struct {
	UserID uint      `json:"user_id" binding:"required"`
	BookID uint      `json:"book_id" binding:"required"`
	DueAt  time.Time `json:"due_at" binding:"required"`
}

type LoanResponseDTO struct {
	ID         uint       `json:"id"`
	UserID     uint       `json:"user_id"`
	BookID     uint       `json:"book_id"`
	BorrowedAt time.Time  `json:"borrowed_at"`
	DueAt      time.Time  `json:"due_at"`
	ReturnedAt *time.Time `json:"returned_at"`
	FineCents  int64      `json:"fine_cents"`
}

type LoanService interface {
	LoanBook(dto *LoanRequestDTO) (*LoanResponseDTO, error)
	ReturnBook(loanID uint) (*LoanResponseDTO, error)
}

type LoanServiceImpl struct {
	loanDAO daos.LoanDAO
	bookDAO daos.BookDAO
	userDAO daos.UserDAO
}

func NewLoanService() *LoanServiceImpl {
	return &LoanServiceImpl{
		loanDAO: daos.NewLoanDAO(),
		bookDAO: daos.NewBookDAO(),
		userDAO: daos.NewUserDAO(),
	}
}

func (s *LoanServiceImpl) LoanBook(dto *LoanRequestDTO) (*LoanResponseDTO, error) {
	// Begin transaction
	returnValue := &LoanResponseDTO{}
	err := s.bookDAO.GetDB().Transaction(func(tx *gorm.DB) error {

		user, err := s.userDAO.FindUserByID(dto.UserID)
		if err != nil {
			return errors.New("user not found")
		}

		book, err := s.bookDAO.GetBookByID(dto.BookID)
		if err != nil {
			return errors.New("book not found")
		}

		if book.Available <= 0 {
			return errors.New("book not available")
		}

		loan := &models.Loan{
			UserID:     user.ID,
			BookID:     book.ID,
			BorrowedAt: time.Now(),
			DueAt:      dto.DueAt,
		}

		if err := s.loanDAO.CreateLoan(loan, tx); err != nil {
			return err
		}

		book.Available -= 1
		if err := s.bookDAO.UpdateBook(book); err != nil {
			return err
		}

		*returnValue = LoanResponseDTO{
			ID:         loan.ID,
			UserID:     loan.UserID,
			BookID:     loan.BookID,
			BorrowedAt: loan.BorrowedAt,
			DueAt:      loan.DueAt,
		}
		return nil
	})
	return returnValue, err
}

func (s *LoanServiceImpl) ReturnBook(loanID uint) (*LoanResponseDTO, error) {
	loan, err := s.loanDAO.FindLoanByID(loanID)
	if err != nil {
		return nil, err
	}

	if loan.ReturnedAt != nil {
		return nil, errors.New("book already returned")
	}

	book, err := s.bookDAO.GetBookByID(loan.BookID)
	if err != nil {
		return nil, err
	}

	// eg: 100 cents per day late
	fineCents := int64(0)
	if time.Now().After(loan.DueAt) {
		daysLate := int(time.Since(loan.DueAt).Hours() / 24)
		if daysLate > 0 {
			fineCents = int64(daysLate * 100)
		}
	}

	if err := s.loanDAO.MarkLoanAsReturned(loan, fineCents); err != nil {
		return nil, err
	}

	book.Available += 1
	if err := s.bookDAO.UpdateBook(book); err != nil {
		return nil, err
	}

	return &LoanResponseDTO{
		ID:         loan.ID,
		UserID:     loan.UserID,
		BookID:     loan.BookID,
		BorrowedAt: loan.BorrowedAt,
		DueAt:      loan.DueAt,
		ReturnedAt: loan.ReturnedAt,
		FineCents:  fineCents,
	}, nil
}
