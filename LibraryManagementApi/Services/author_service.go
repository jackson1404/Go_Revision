package services

import (
	daos "jackson.com/libraryapisystem/DAOs"
	"jackson.com/libraryapisystem/models"
)

type AuthorRequestDTO struct {
	Name string `json:"name" binding:"required"`
	Bio  string `json:"bio"`
}

type AuthorResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type AuthorService interface {
	CreateAuthor(dto *AuthorRequestDTO) error
	GetAllAuthors() ([]AuthorResponseDTO, error)
	GetAuthorByID(id uint) (*AuthorResponseDTO, error)
	UpdateAuthor(id uint, dto *AuthorRequestDTO) error
	DeleteAuthor(id uint) error
}

type AuthorServiceImpl struct {
	authorDAO daos.AuthorDAO
}

func NewAuthorService() *AuthorServiceImpl {
	return &AuthorServiceImpl{
		authorDAO: daos.NewAuthorDAO(), // internally creates DAO
	}
}

func (s *AuthorServiceImpl) CreateAuthor(dto *AuthorRequestDTO) error {
	author := &models.Author{
		Name: dto.Name,
		Bio:  dto.Bio,
	}
	return s.authorDAO.CreateAuthor(author)
}

func (s *AuthorServiceImpl) GetAllAuthors() ([]AuthorResponseDTO, error) {
	authors, err := s.authorDAO.FindAllAuthors()
	if err != nil {
		return nil, err
	}
	var result []AuthorResponseDTO
	for _, a := range authors {
		result = append(result, AuthorResponseDTO{
			ID:   a.ID,
			Name: a.Name,
			Bio:  a.Bio,
		})
	}
	return result, nil
}

func (s *AuthorServiceImpl) GetAuthorByID(id uint) (*AuthorResponseDTO, error) {
	author, err := s.authorDAO.FindAuthorByID(id)
	if err != nil {
		return nil, err
	}
	return &AuthorResponseDTO{
		ID:   author.ID,
		Name: author.Name,
		Bio:  author.Bio,
	}, nil
}

func (s *AuthorServiceImpl) UpdateAuthor(id uint, dto *AuthorRequestDTO) error {
	author, err := s.authorDAO.FindAuthorByID(id)
	if err != nil {
		return err
	}
	author.Name = dto.Name
	author.Bio = dto.Bio
	return s.authorDAO.UpdateAuthor(author)
}

func (s *AuthorServiceImpl) DeleteAuthor(id uint) error {
	return s.authorDAO.DeleteAuthor(id)
}
