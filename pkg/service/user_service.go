package service

import (
	"context"
	"errors"

	"github.com/bitwyre/template-golang/pkg/datastore/mysql/entity"
	"github.com/bitwyre/template-golang/pkg/repository"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type IUserService interface {
	GetUser(id int, c context.Context) (*entity.User, error)
	GetUser2(id int, c context.Context) (string, error)
}

type userService struct {
	repository *repository.Repository
}

func NewUserService(repository *repository.Repository) *userService {
	return &userService{repository: repository}
}

func (s userService) GetUser(id int, c context.Context) (*entity.User, error) {
	var data = entity.User{}
	// Use the global TracerProvider.
	tr := otel.Tracer("GetUserService")
	_, span := tr.Start(c, "GetUserService")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	data, err := s.repository.UserRepo.FindById(id, c)
	if err != nil {
		return &data, errors.New("user not found")
	}

	return &data, nil
}

func (s userService) GetUser2(id int, c context.Context) (string, error) {
	return "OK", nil
}
