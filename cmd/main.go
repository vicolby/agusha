package main

import (
	"log"

	"github.com/gimmefear/dswv3/internal/domain"
	"github.com/gimmefear/dswv3/internal/handler"
	"github.com/gimmefear/dswv3/internal/repository"
	"github.com/gimmefear/dswv3/internal/services"
	"github.com/gimmefear/dswv3/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	workspaceRepo := repository.NewWorkspaceRepository(db)
	workspaceService := services.NewWorkspaceService(workspaceRepo)

	handler.NewWorkspaceHandler(e, workspaceService)

	log.Fatal(e.Start(":1234"))
}
