package grpc

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bitwyre/template-golang/pkg/service"
)

type GetUserRPC interface {
	GetUser(ctx context.Context, payload *UserPayload) (*UserDataResp, error)
}

type GetUserRPCService struct {
	service *service.Service
}

func NewGetUserRPC(service *service.Service) *GetUserRPCService {
	return &GetUserRPCService{
		service: service,
	}
}

func (s *GetUserRPCService) GetUser(ctx context.Context, payload *UserPayload) (*UserDataResp, error) {
	userId, err := strconv.Atoi(payload.Id)
	if err != nil {
		return nil, err
	}

	fmt.Println("Got Request from Client")

	userData, err := s.service.IUserService.GetUser(userId, ctx)

	return &UserDataResp{
		Email:    userData.Email,
		Status:   string(rune(userData.Status)),
		UserCode: userData.UserCode,
	}, nil

}
