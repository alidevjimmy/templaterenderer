package service

import (
	"github.com/alidevjimmy/templaterenderer/internal/entity"
	"github.com/alidevjimmy/templaterenderer/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Join(user *entity.User) error {
	return u.userRepository.Save(user.Username, user)
}

func (u *UserService) ProfileByUsername(username string) (*entity.User, error) {
	return u.userRepository.Fetch(username)
}
