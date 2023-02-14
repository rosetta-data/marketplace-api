package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/evgeniy-dammer/emenu-api/internal/config"
	"github.com/evgeniy-dammer/emenu-api/internal/handler"
	"github.com/evgeniy-dammer/emenu-api/internal/model"
	"github.com/evgeniy-dammer/emenu-api/internal/repository"
	"github.com/evgeniy-dammer/emenu-api/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// set log format as JSON
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// initializing config
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	// loading env variables
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// establishing database connection
	database, err := repository.NewPostgresDB(
		model.DBConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			Username: viper.GetString("database.username"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   viper.GetString("database.dbname"),
			SSLMode:  viper.GetString("database.sslmode"),
		},
	)
	if err != nil {
		logrus.Fatalf("failed to initialize database: %s", err)
	}

	// dependency injections
	repos := repository.NewRepository(database)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// create new server
	srv := new(model.Server)

	go func() {
		// run server
		if err = srv.Run(viper.GetString("application.port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Println("Application started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Application shutdown...")

	if err = srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error ocured on server shutdown: %s", err.Error())
	}
}
