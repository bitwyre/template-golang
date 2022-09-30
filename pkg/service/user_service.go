package service

import (
	"errors"

	"github.com/bitwyre/template-golang/pkg/datastore/postgres/entity"
	"github.com/bitwyre/template-golang/pkg/repository"
)

type IUserService interface {
	GetUser(id int) (*entity.User, error)
}

type userService struct {
	repository *repository.Repository
}

func NewUserService(repository *repository.Repository) *userService {
	return &userService{repository: repository}
}

func (s userService) GetUser(id int) (*entity.User, error) {
	var data = entity.User{}

	data, err := s.repository.UserRepo.FindById(id)
	if err != nil {
		return &data, errors.New("user not found")
	}

	return &data, nil
}
