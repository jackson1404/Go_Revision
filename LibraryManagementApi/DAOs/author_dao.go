package daos

import (
	"gorm.io/gorm"
	"jackson.com/libraryapisystem/Configurations/db_configs"
	"jackson.com/libraryapisystem/models"
)

type AuthorDAO interface {
	CreateAuthor(author *models.Author) error
	FindAllAuthors() ([]models.Author, error)
	FindAuthorByID(id uint) (*models.Author, error)
	UpdateAuthor(author *models.Author) error
	DeleteAuthor(id uint) error
}

type AuthorDAOImpl struct {
	db *gorm.DB
}

func NewAuthorDAO() *AuthorDAOImpl {
	return &AuthorDAOImpl{
		db: db_configs.DB,
	}
}

func (dao *AuthorDAOImpl) CreateAuthor(author *models.Author) error {
	return dao.db.Create(author).Error
}

func (dao *AuthorDAOImpl) FindAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	err := dao.db.Find(&authors).Error
	return authors, err
}

func (dao *AuthorDAOImpl) FindAuthorByID(id uint) (*models.Author, error) {
	var author models.Author
	err := dao.db.First(&author, id).Error
	return &author, err
}

func (dao *AuthorDAOImpl) UpdateAuthor(author *models.Author) error {
	return dao.db.Save(author).Error
}

func (dao *AuthorDAOImpl) DeleteAuthor(id uint) error {
	return dao.db.Delete(&models.Author{}, id).Error
}
