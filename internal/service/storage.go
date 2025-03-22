package service

import "github.com/vrabber/storage/internal/repository"

type Service interface{}

type Implementation struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &Implementation{repo: repo}
}
