package daos

import (
	"gorm.io/gorm"
	"jackson.com/libraryapisystem/Configurations/db_configs"
	"jackson.com/libraryapisystem/models"
)

type CategoryDAO interface {
	CreateCategory(category *models.Category) error
	FindAllCategories() ([]models.Category, error)
	FindCategoryById(id uint) (*models.Category, error)
	UpdateCategory(category *models.Category) error
	DeleteCategory(id uint) error
}

type CategoryDAOImpl struct {
	dbClient *gorm.DB
}

func NewCategoryDAO() *CategoryDAOImpl {
	return &CategoryDAOImpl{
		dbClient: db_configs.DB,
	}
}

func (dao *CategoryDAOImpl) CreateCategory(category *models.Category) error {
	return dao.dbClient.Create(category).Error
}

func (dao *CategoryDAOImpl) FindAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := dao.dbClient.Order("created_at desc").Find(&categories).Error
	return categories, err
}

func (dao *CategoryDAOImpl) FindCategoryById(id uint) (*models.Category, error) {
	var category models.Category

	err := dao.dbClient.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (dao *CategoryDAOImpl) UpdateCategory(category *models.Category) error {
	return dao.dbClient.Save(category).Error
}

func (dao *CategoryDAOImpl) DeleteCategory(id uint) error {
	return dao.dbClient.Delete(&models.Category{}, id).Error
}
