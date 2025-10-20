package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	daos "jackson.com/libraryapisystem/DAOs"
	"jackson.com/libraryapisystem/models"
)

type UserRequestDTO struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"` // plaintext here, hash in service
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

type UserResponseDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

type UserService interface {
	CreateUser(dto *UserRequestDTO) error
	GetAllUsers() ([]UserResponseDTO, error)
	GetUserByID(id uint) (*UserResponseDTO, error)
	UpdateUser(id uint, dto *UserRequestDTO) error
	DeleteUser(id uint) error
}

type UserServiceImpl struct {
	userDAO daos.UserDAO
}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{
		userDAO: daos.NewUserDAO(),
	}
}

// CREATE
func (s *UserServiceImpl) CreateUser(dto *UserRequestDTO) error {
	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:     dto.Username,
		Email:        dto.Email,
		PasswordHash: string(hash),
		FullName:     dto.FullName,
		Phone:        dto.Phone,
	}
	return s.userDAO.CreateUser(user)
}

// GET ALL
func (s *UserServiceImpl) GetAllUsers() ([]UserResponseDTO, error) {
	users, err := s.userDAO.FindAllUsers()
	if err != nil {
		return nil, err
	}
	var result []UserResponseDTO
	for _, u := range users {
		result = append(result, UserResponseDTO{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
			FullName: u.FullName,
			Phone:    u.Phone,
		})
	}
	return result, nil
}

// GET BY ID
func (s *UserServiceImpl) GetUserByID(id uint) (*UserResponseDTO, error) {
	user, err := s.userDAO.FindUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &UserResponseDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
		Phone:    user.Phone,
	}, nil
}

// UPDATE
func (s *UserServiceImpl) UpdateUser(id uint, dto *UserRequestDTO) error {
	user, err := s.userDAO.FindUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	user.Username = dto.Username
	user.Email = dto.Email
	user.FullName = dto.FullName
	user.Phone = dto.Phone

	// Update password only if provided
	if dto.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.PasswordHash = string(hash)
	}

	return s.userDAO.UpdateUser(user)
}

// DELETE
func (s *UserServiceImpl) DeleteUser(id uint) error {
	return s.userDAO.DeleteUser(id)
}
