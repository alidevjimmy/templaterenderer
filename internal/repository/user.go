package repository

import (
	"fmt"
	"github.com/alidevjimmy/templaterenderer/internal/entity"
	"sync"
)

var (
	ErrUserNotFound = fmt.Errorf("user not found")
)

type UserRepository struct {
	users map[string]*entity.User
	mu    sync.RWMutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*entity.User),
	}
}

// Save creates user if not exists and updates user if exists
func (u *UserRepository) Save(key string, data *entity.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.users[key] = data
	return nil
}

func (u *UserRepository) Fetch(key string) (*entity.User, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	user, ok := u.users[key]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}
