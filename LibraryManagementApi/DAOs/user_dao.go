package daos

import (
	"gorm.io/gorm"
	"jackson.com/libraryapisystem/Configurations/db_configs"
	"jackson.com/libraryapisystem/models"
)

type UserDAO interface {
	CreateUser(user *models.User) error
	FindAllUsers() ([]models.User, error)
	FindUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type UserDAOImpl struct {
	db *gorm.DB
}

func NewUserDAO() *UserDAOImpl {
	return &UserDAOImpl{db: db_configs.DB}
}

func (dao *UserDAOImpl) CreateUser(user *models.User) error {
	return dao.db.Create(user).Error
}

func (dao *UserDAOImpl) FindAllUsers() ([]models.User, error) {
	var users []models.User
	err := dao.db.Find(&users).Error
	return users, err
}

func (dao *UserDAOImpl) FindUserByID(id uint) (*models.User, error) {
	var user models.User
	err := dao.db.First(&user, id).Error
	return &user, err
}

func (dao *UserDAOImpl) UpdateUser(user *models.User) error {
	return dao.db.Save(user).Error
}

func (dao *UserDAOImpl) DeleteUser(id uint) error {
	return dao.db.Delete(&models.User{}, id).Error
}
