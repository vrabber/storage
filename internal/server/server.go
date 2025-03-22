package server

import (
	"errors"

	pb "github.com/vrabber/storage/gen/storage"
	"github.com/vrabber/storage/internal/service"
)

type Server struct {
	pb.UnimplementedStorageServiceServer
	srv service.Service
}

func NewServer(srv service.Service) *Server {
	return &Server{srv: srv}
}

func (*Server) Run() error {
	return errors.New("unimplemented")
}
