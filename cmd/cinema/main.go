package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ablyamitov/cinema/internal/app/config"
	"github.com/Ablyamitov/cinema/internal/app/db"
	"github.com/Ablyamitov/cinema/internal/app/db/models"
	"github.com/Ablyamitov/cinema/internal/app/server"
	"github.com/Ablyamitov/cinema/internal/app/service"
)


func main() {
	conf, err := setupConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	gormDB := db.Connect(conf.DB.URL)

	if err := gormDB.AutoMigrate(&models.Movie{}, &models.Comment{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	movieService := service.NewMovieService(gormDB)

	app := initializeServer(conf, movieService)
	app.Run()

	waitForShutdown(app)
}

func setupConfig() (*config.Config, error) {
	return config.MustLoad()
}

func initializeServer(conf *config.Config, movieService service.MovieService) server.App {
	return server.NewServer(conf.Server.Host, conf.Server.Port, &movieService)
}

func waitForShutdown(app server.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	app.Stop(ctx)
}
