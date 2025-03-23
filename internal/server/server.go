package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	pb "github.com/vrabber/storage/gen/storage"
	"github.com/vrabber/storage/internal/config"
	"github.com/vrabber/storage/internal/service"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedStorageServiceServer
	srv  service.Service
	conf config.ServerConfig
}

func NewServer(srv service.Service, conf config.ServerConfig) *Server {
	return &Server{
		srv:  srv,
		conf: conf,
	}
}

func (s *Server) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.conf.Host, s.conf.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	go func() {
		<-ctx.Done()
		grpcServer.GracefulStop()
	}()

	pb.RegisterStorageServiceServer(grpcServer, s)

	slog.Info("starting gRPC server", "host", s.conf.Host, "port", s.conf.Port)
	return grpcServer.Serve(lis)
}
