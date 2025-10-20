package services

import (
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"
	daos "jackson.com/libraryapisystem/DAOs"
	"jackson.com/libraryapisystem/models"
)

type BookRequestDTO struct {
	Title       string `json:"title" binding:"required"`
	ISBN        string `json:"isbn" binding:"required"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id" binding:"required"`
	AuthorIDs   []uint `json:"author_ids" binding:"required"`
	TotalCopies uint   `json:"total_copies"`
}

type BookResponseDTO struct {
	ID          uint        `json:"id"`
	Title       string      `json:"title"`
	ISBN        string      `json:"isbn"`
	Description string      `json:"description"`
	Category    CategoryDTO `json:"category"`
	Authors     []AuthorDTO `json:"authors"`
	TotalCopies uint        `json:"total_copies"`
	Available   int         `json:"available"`
}

type CategoryDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type AuthorDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type BookService interface {
	RegisterBook(dto *BookRequestDTO) error
	GetAllBooks() ([]BookResponseDTO, error)
	GetBookByID(id uint) (*BookResponseDTO, error)
	UpdateBook(id uint, dto *BookRequestDTO) error
	DeleteBook(id uint) error
}

type BookServiceImpl struct {
	bookDAO daos.BookDAO
}

// Shortcut constructor
func NewBookService() *BookServiceImpl {
	return &BookServiceImpl{
		bookDAO: daos.NewBookDAO(),
	}
}

func (s *BookServiceImpl) RegisterBook(dto *BookRequestDTO) error {
	dbClient := s.bookDAO.(*daos.BookDAOImpl).GetDB()

	return dbClient.Transaction(func(tx *gorm.DB) error {

		var category models.Category
		if err := tx.First(&category, dto.CategoryID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("category not found")
			}
			return err
		}

		var authors []models.Author
		for _, authorID := range dto.AuthorIDs {
			var author models.Author
			if err := tx.First(&author, authorID).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return errors.New("author not found with ID: " + strconv.Itoa(int(authorID)))
				}
				return err
			}
			authors = append(authors, author)
		}

		book := &models.Book{
			Title:       dto.Title,
			ISBN:        dto.ISBN,
			Description: dto.Description,
			CategoryID:  dto.CategoryID,
			TotalCopies: dto.TotalCopies,
			Available:   int(dto.TotalCopies),
		}

		if err := s.bookDAO.Create(book, tx); err != nil {
			return err // rollback
		}

		// 4️⃣ Attach Authors
		if err := s.bookDAO.ReplaceAuthors(book, authors, tx); err != nil {
			return err // rollback
		}

		return nil
	})
}

// func (s *BookServiceImpl) RegisterBook(dto *BookRequestDTO) error {
// 	book := &models.Book{
// 		Title:       dto.Title,
// 		ISBN:        dto.ISBN,
// 		Description: dto.Description,
// 		CategoryID:  dto.CategoryID,
// 		TotalCopies: dto.TotalCopies,
// 		Available:   int(dto.TotalCopies),
// 	}

// 	var authors []models.Author
// 	for _, id := range dto.AuthorIDs {
// 		authors = append(authors, models.Author{Model: gorm.Model{ID: id}})
// 	}

// 	return s.bookDAO.CreateBookWithAuthors(book, authors)
// }

// GET ALL
func (s *BookServiceImpl) GetAllBooks() ([]BookResponseDTO, error) {
	books, err := s.bookDAO.GetAllBooks()
	if err != nil {
		return nil, err
	}

	var result []BookResponseDTO
	for _, b := range books {
		var authors []AuthorDTO
		for _, a := range b.Authors {
			authors = append(authors, AuthorDTO{ID: a.ID, Name: a.Name})
		}

		result = append(result, BookResponseDTO{
			ID:          b.ID,
			Title:       b.Title,
			ISBN:        b.ISBN,
			Description: b.Description,
			Category:    CategoryDTO{ID: b.Category.ID, Name: b.Category.Name},
			Authors:     authors,
			TotalCopies: b.TotalCopies,
			Available:   b.Available,
		})
	}
	return result, nil
}

// GET BY ID
func (s *BookServiceImpl) GetBookByID(id uint) (*BookResponseDTO, error) {
	book, err := s.bookDAO.GetBookByID(id)
	if err != nil {
		return nil, errors.New("book not found")
	}

	var authors []AuthorDTO
	for _, a := range book.Authors {
		authors = append(authors, AuthorDTO{ID: a.ID, Name: a.Name})
	}

	return &BookResponseDTO{
		ID:          book.ID,
		Title:       book.Title,
		ISBN:        book.ISBN,
		Description: book.Description,
		Category:    CategoryDTO{ID: book.Category.ID, Name: book.Category.Name},
		Authors:     authors,
		TotalCopies: book.TotalCopies,
		Available:   book.Available,
	}, nil
}

func (s *BookServiceImpl) UpdateBook(id uint, dto *BookRequestDTO) error {
	// 1️⃣ Get the book
	book, err := s.bookDAO.GetBookByID(id)
	if err != nil {
		return errors.New("book not found")
	}

	// 2️⃣ Validate category
	exists, err := s.bookDAO.CategoryExists(dto.CategoryID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("category not found")
	}
	book.CategoryID = dto.CategoryID

	// 3️⃣ Validate authors
	authors, err := s.bookDAO.AuthorsExist(dto.AuthorIDs)
	if err != nil {
		return err
	}
	book.Authors = authors

	// 4️⃣ Update other fields
	book.Title = dto.Title
	book.ISBN = dto.ISBN
	book.Description = dto.Description
	book.TotalCopies = dto.TotalCopies
	book.Available = int(dto.TotalCopies)

	fmt.Print("Book id for update:", book.CategoryID)
	// 5️⃣ Save book (DAO handles authors join table)
	return s.bookDAO.UpdateBook(book)
}

// DELETE
func (s *BookServiceImpl) DeleteBook(id uint) error {
	return s.bookDAO.DeleteBook(id)
}
