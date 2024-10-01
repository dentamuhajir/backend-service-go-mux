package service

import (
	"github.com/dentamuhajir/backend-service-go-mysql/internal/model"
	"github.com/dentamuhajir/backend-service-go-mysql/internal/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{
		userRepository: r,
	}
}

func (s UserService) GetListUser() ([]model.User, error) {
	return s.userRepository.GetUser()
}