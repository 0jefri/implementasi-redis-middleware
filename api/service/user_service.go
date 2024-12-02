package service

import (
	"github.com/project-sistem-voucher/api/model"
	"github.com/project-sistem-voucher/api/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterNewUser(payload *model.User) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(usrRepo repository.UserRepository) UserService {
	return &userService{repo: usrRepo}
}

func (s *userService) RegisterNewUser(payload *model.User) (*model.User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	password := string(bytes)

	payload.Password = password

	user, err := s.repo.Create(payload)

	userResponse := model.User{
		// ID: user.ID,
		Username: user.Username,
		Password: user.Password,
	}

	return &userResponse, err
}

func (s *userService) FindByUsername(username string) (*model.User, error) {
	return s.repo.GetUsername(username)
}
