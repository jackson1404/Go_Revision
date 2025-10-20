package daos

import (
	"errors"
	"strconv"

	"gorm.io/gorm"
	"jackson.com/libraryapisystem/Configurations/db_configs"
	"jackson.com/libraryapisystem/models"
)

type BookDAO interface {
	Create(book *models.Book, tx *gorm.DB) error
	GetAllBooks() ([]models.Book, error)
	ReplaceAuthors(book *models.Book, authors []models.Author, tx *gorm.DB) error
	GetBookByID(id uint) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
	GetDB() *gorm.DB
	AuthorsExist(authorIDs []uint) ([]models.Author, error)
	CategoryExists(categoryID uint) (bool, error)
}

type BookDAOImpl struct {
	db *gorm.DB
}

// Shortcut constructor
func NewBookDAO() *BookDAOImpl {
	return &BookDAOImpl{db: db_configs.DB}
}

func (dao *BookDAOImpl) GetDB() *gorm.DB {
	return dao.db
}

func (dao *BookDAOImpl) Create(book *models.Book, tx *gorm.DB) error {
	return tx.Create(book).Error
}

func (dao *BookDAOImpl) ReplaceAuthors(book *models.Book, authors []models.Author, tx *gorm.DB) error {
	return tx.Model(book).Association("Authors").Replace(authors)
}

// GET ALL
func (dao *BookDAOImpl) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := dao.db.Preload("Category").Preload("Authors").Find(&books).Error
	return books, err
}

// GET BY ID
func (dao *BookDAOImpl) GetBookByID(id uint) (*models.Book, error) {
	var book models.Book
	err := dao.db.Preload("Category").Preload("Authors").First(&book, id).Error
	return &book, err
}

func (dao *BookDAOImpl) UpdateBook(book *models.Book) error {

	return dao.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Book{}).Where("id = ?", book.ID).Updates(map[string]any{
			"title":        book.Title,
			"isbn":         book.ISBN,
			"description":  book.Description,
			"category_id":  book.CategoryID,
			"total_copies": book.TotalCopies,
			"available":    book.Available,
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(book).Association("Authors").Replace(book.Authors); err != nil {
			return err
		}

		return nil
	})
}

// DELETE
func (dao *BookDAOImpl) DeleteBook(id uint) error {
	return dao.db.Delete(&models.Book{}, id).Error
}

func (dao *BookDAOImpl) CategoryExists(categoryID uint) (bool, error) {
	var category models.Category
	if err := dao.db.First(&category, categoryID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (dao *BookDAOImpl) AuthorsExist(authorIDs []uint) ([]models.Author, error) {
	var authors []models.Author
	for _, id := range authorIDs {
		var author models.Author
		if err := dao.db.First(&author, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("author not found with ID: " + strconv.Itoa(int(id)))
			}
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}
