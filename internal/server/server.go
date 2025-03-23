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
	ctx  context.Context
	srv  service.Service
	conf config.ServerConfig
}

func NewServer(ctx context.Context, srv service.Service, conf config.ServerConfig) *Server {
	return &Server{
		ctx:  ctx,
		srv:  srv,
		conf: conf,
	}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.conf.Host, s.conf.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	go func() {
		<-s.ctx.Done()
		grpcServer.GracefulStop()
	}()

	pb.RegisterStorageServiceServer(grpcServer, s)

	slog.Info("starting gRPC server", "host", s.conf.Host, "port", s.conf.Port)
	return grpcServer.Serve(lis)
}
