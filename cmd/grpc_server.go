package cmd

import (
	"log"
	"net"

	"github.com/bitwyre/template-golang/pkg/lib"
	"google.golang.org/grpc"
)

func StartGRPCServer(grpcServer *grpc.Server) {
	listener, err := net.Listen("tcp", ":"+lib.AppConfig.App.GrpcServerPort)
	if err != nil {
		log.Fatalf("🔴 Couldn't start Listener %v", err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("🔴 Couldn't connect to gRPC Server %v", err)
	}
}
