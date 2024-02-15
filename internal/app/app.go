package app

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/mustthink/microblog/internal/config"
	"github.com/mustthink/microblog/internal/service"
	"github.com/mustthink/microblog/internal/storage"
)

type App struct {
	config *config.Config
	server *grpc.Server
	db     *storage.Storage
	logger *logrus.Logger
}

func New() *App {
	cfg := config.MustLoad()

	logger := logrus.New()
	logger.SetLevel(cfg.GetLogLevel())

	db, err := storage.New(cfg.Database)
	if err != nil {
		logger.Fatalf("couldn't create new storage w err %s", err.Error())
	}

	server := grpc.NewServer()

	microblogService := service.NewMicroblog(db, db, db, logger.WithField("service", "microblog"))
	microblogService.Register(server)

	return &App{
		config: cfg,
		db:     db,
		server: server,
		logger: logger,
	}
}

// Run runs gRPC server
func (a *App) Run() error {
	const op = "app.Run"
	log := a.logger.WithField("op", op)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.config.GRPC.Port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr", listener.Addr().String()))
	if err := a.server.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// MustRun runs application and panics if an error occurs
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Stop stops application gracefully
func (a *App) Stop() {
	const op = "app.Stop"
	log := a.logger.WithField("op", op)
	log.Info("stopping app")

	a.server.GracefulStop()
	log.Debug("gRPC server stopped")

	if err := a.db.Close(); err != nil {
		log.Errorf("couldn't close db w err %s", err.Error())
	}
	log.Debug("db connection closed")

	log.Info("app stopped")
}
