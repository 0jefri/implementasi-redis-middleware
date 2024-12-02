package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/project-sistem-voucher/api/repository"
	"github.com/project-sistem-voucher/config"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (string, error)
}

type authService struct {
	userRepo repository.UserRepository
	redis    config.Cacher
}

func NewAuthService(userRepo repository.UserRepository, redisClient config.Cacher) AuthService {
	return &authService{
		userRepo: userRepo,
		redis:    redisClient,
	}
}

func (s *authService) Login(username, password string) (string, error) {
	// ctx := context.Background()

	// Find user by username
	user, err := s.userRepo.GetUsername(username)
	if err != nil || user == nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token := uuid.NewString()
	idKey := user.Username

	// sessionData := map[string]interface{}{
	// 	"userID":   user.ID,
	// 	"username": user.Username,
	// }
	err = s.redis.Set(idKey, token)
	if err != nil {
		return "", err
	}

	// s.redis.Expire(ctx, token, time.Hour)

	return token, nil
}
