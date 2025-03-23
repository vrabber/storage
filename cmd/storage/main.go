package main

import (
	"context"
	"flag"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/vrabber/storage/internal/config"
	"github.com/vrabber/storage/internal/db"
	"github.com/vrabber/storage/internal/repository"
	"github.com/vrabber/storage/internal/server"
	"github.com/vrabber/storage/internal/service"
	"github.com/vrabber/storage/internal/store"
	"github.com/vrabber/storage/internal/store/driver"
)

var configSource string

func init() {
	flag.StringVar(&configSource, "config-source", "env", "config source, supported values are: env, yaml")
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	cnf, err := config.Load(configSource)
	if err != nil {
		slog.Error("failed to load config", "err", err)
		panic(err)
	}

	pool, err := db.CreatePool(ctx, cnf.Database)
	if err != nil {
		slog.Error("failed to create pool", "err", err)
		panic(err)
	}
	defer pool.Close()

	repo := repository.NewRepositoryImplementation(pool)
	fileStore := store.NewImplementation()

	if err = setupStoreDrivers(fileStore); err != nil {
		slog.Error("failed to setup store drivers", "err", err)
		panic(err)
	}

	srv := service.NewService(repo)

	server_ := server.NewServer(srv)
	if err = server_.Run(); err != nil {
		slog.Error("application stopped", "err", err)
	}
}

func setupStoreDrivers(s store.Store) error {
	if err := s.RegisterDriver(driver.NewLocalDriver(".")); err != nil {
		return err
	}
	return nil
}
