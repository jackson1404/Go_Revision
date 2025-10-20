package services

import (
	"errors"

	daos "jackson.com/libraryapisystem/DAOs"
	models "jackson.com/libraryapisystem/models"
)

type CategoryRequestDTO struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}

type CategoryResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CategoryService interface {
	CreateCategory(dto *CategoryRequestDTO) error
	FindAllCategories() ([]CategoryResponseDTO, error)
	FindCategoryById(id uint) (*CategoryResponseDTO, error)
	UpdateCategory(id uint, dto *CategoryRequestDTO) error
	DeleteCategory(id uint) error
}

type CategoryServiceImpl struct {
	categoryDAO daos.CategoryDAO
}

func NewCategoryService() *CategoryServiceImpl {
	return &CategoryServiceImpl{
		categoryDAO: daos.NewCategoryDAO(),
	}
}

// ✅ CREATE
func (service *CategoryServiceImpl) CreateCategory(dto *CategoryRequestDTO) error {
	category := &models.Category{
		Name: dto.Name,
		Slug: dto.Slug,
	}
	return service.categoryDAO.CreateCategory(category)
}

// ✅ READ ALL
func (service *CategoryServiceImpl) FindAllCategories() ([]CategoryResponseDTO, error) {
	categories, err := service.categoryDAO.FindAllCategories()
	if err != nil {
		return nil, err
	}

	var result []CategoryResponseDTO
	for _, cat := range categories {
		result = append(result, CategoryResponseDTO{
			ID:   cat.ID,
			Name: cat.Name,
			Slug: cat.Slug,
		})
	}
	return result, nil
}

// ✅ READ ONE
func (service *CategoryServiceImpl) FindCategoryById(id uint) (*CategoryResponseDTO, error) {
	category, err := service.categoryDAO.FindCategoryById(id)
	if err != nil {
		return nil, errors.New("category not found")
	}

	return &CategoryResponseDTO{
		ID:   category.ID,
		Name: category.Name,
		Slug: category.Slug,
	}, nil
}

// ✅ UPDATE
func (service *CategoryServiceImpl) UpdateCategory(id uint, dto *CategoryRequestDTO) error {
	category, err := service.categoryDAO.FindCategoryById(id)
	if err != nil {
		return errors.New("category not found")
	}

	category.Name = dto.Name
	category.Slug = dto.Slug

	return service.categoryDAO.UpdateCategory(category)
}

// ✅ DELETE
func (service *CategoryServiceImpl) DeleteCategory(id uint) error {
	return service.categoryDAO.DeleteCategory(id)
}
