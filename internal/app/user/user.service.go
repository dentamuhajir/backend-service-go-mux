package user

import (
	"github.com/dentamuhajir/backend-service-go-mysql/internal/app/user"
	"github.com/dentamuhajir/backend-service-go-mysql/internal/model"
)

type UserService struct {
	userRepository user.UserRepository
}

func NewUserService(r user.UserRepository) *UserService {
	return &UserService{
		userRepository: r,
	}
}

func (s UserService) GetListUser() ([]model.User, error) {
	return s.userRepository.GetUser()
}
