package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sigit14ap/user-service/config"
	delivery "github.com/sigit14ap/user-service/internal/delivery/http"
	"github.com/sigit14ap/user-service/internal/domain"
	repository "github.com/sigit14ap/user-service/internal/repository/mysql"
	"github.com/sigit14ap/user-service/internal/router"
	"github.com/sigit14ap/user-service/internal/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("failed to auto-migrate User model: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := usecase.NewUserUsecase(userRepo)
	userHandler := delivery.NewUserHandler(userService)

	router := router.NewRouter(userHandler)

	log.Fatal(router.Run(":" + os.Getenv("APP_PORT")))
}
