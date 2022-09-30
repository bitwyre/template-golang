package service

import "github.com/bitwyre/template-golang/pkg/repository"

type Service struct {
	IUserService IUserService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		IUserService: NewUserService(repository),
	}
}
