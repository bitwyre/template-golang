package grpc

import (
	"github.com/bitwyre/template-golang/pkg/service"
	"google.golang.org/grpc"
)

type GPRCService struct {
	GetUserRPC GetUserRPC
}

func NewGRPC(service *service.Service) *GPRCService {
	return &GPRCService{
		GetUserRPC: NewGetUserRPC(service),
	}
}

func RegisterGRPCService(g *grpc.Server, s *GPRCService) {
	RegisterGetUserServiceServer(g, s.GetUserRPC)
}
