package services

import (
	"context"
	"go-hexagonal-template/internal/core/domain"
	"go-hexagonal-template/internal/core/ports"
)

type UserService struct {
	UserRepo ports.IUserRepository
}

func (u UserService) InsertUser(ctx context.Context, user domain.User) error {
	// business logic can be added here if needed
	return u.UserRepo.InsertUser(ctx, user)
}

func (u UserService) GetUsers(ctx context.Context) ([]domain.User, error) {
	// business logic can be added here if needed
	return u.UserRepo.GetUsers(ctx)
}

func NewUserService(repo ports.IUserRepository) ports.IUserService {
	return UserService{
		UserRepo: repo,
	}
}
