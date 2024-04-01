package main

import (
	"log"

	"github.com/gimmefear/dswv3/infra/logging"
	"github.com/gimmefear/dswv3/internal/domain"
	"github.com/gimmefear/dswv3/internal/handler"
	"github.com/gimmefear/dswv3/internal/repository"
	"github.com/gimmefear/dswv3/internal/services"
	"github.com/gimmefear/dswv3/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&domain.Image{}, &domain.User{}, &domain.Workspace{})

	logger := logging.NewZapLogger()

	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("Method", v.Method),
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))
	e.Use(middleware.Recover())

	k8sService, _ := services.NewK8sService()
	workspaceRepo := repository.NewWorkspaceRepository(db)
	workspaceService := services.NewWorkspaceService(workspaceRepo, *k8sService, logger)

	handler.NewWorkspaceHandler(e, workspaceService)

	log.Fatal(e.Start(":1234"))
}
