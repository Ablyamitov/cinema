package server

import (
	"context"
	"fmt"
	"github.com/Ablyamitov/cinema/internal/app/dto"
	"github.com/Ablyamitov/cinema/internal/app/mapper"
	"github.com/Ablyamitov/cinema/internal/app/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"time"
)

type App interface {
	Run()
	Stop(ctx context.Context)
}

type HttpServer struct {
	app          *fiber.App
	host         string
	port         int
	movieService service.MovieService
}

func NewServer(host string, port int, movieService *service.MovieService) App {
	app := fiber.New()

	server := &HttpServer{
		app:          app,
		host:         host,
		port:         port,
		movieService: *movieService,
	}
	server.registerRoutes()
	return server
}

func (s *HttpServer) registerRoutes() {
	s.app.Post("/movies", s.createMovie)
	s.app.Get("/movies/:id", s.getMovie)
	s.app.Get("/movies", s.getAllMovies)
	s.app.Put("/movies/:id", s.updateMovie)
	s.app.Delete("/movies/:id", s.deleteMovie)
}

func (s *HttpServer) Run() {
	go func() {
		// Запуск Fiber-сервера с использованием хоста и порта
		addr := fmt.Sprintf("%s:%d", s.host, s.port)
		if err := s.app.Listen(addr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
}

func (s *HttpServer) Stop(ctx context.Context) {
	shutdownTimeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	if err := s.app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Failed to gracefully shutdown server: %v", err)
	}
	log.Println("Server stopped gracefully")
}

// Создание фильма
func (s *HttpServer) createMovie(c *fiber.Ctx) error {
	var request dto.MovieDTO
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	movie := mapper.MapMovieDTOToMovie(request)

	if err := s.movieService.CreateMovie(&movie); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create movie"})
	}

	response := mapper.MapMovieToMovieDTO(movie)
	return c.Status(fiber.StatusCreated).JSON(response)
}

// Получение фильма по ID
func (s *HttpServer) getMovie(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	movie, err := s.movieService.GetMovieByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Movie not found"})
	}
	response := mapper.MapMovieToMovieDTO(*movie)
	return c.JSON(response)
}

// Получение всех фильмов
func (s *HttpServer) getAllMovies(c *fiber.Ctx) error {
	movies, err := s.movieService.GetAllMovies()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch movies"})
	}
	response := mapper.MapMoviesToMoviesDto(movies)
	return c.JSON(response)
}

// Обновление фильма
func (s *HttpServer) updateMovie(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	movie, err := s.movieService.GetMovieByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Movie not found"})
	}

	response := mapper.MapMovieToMovieDTO(*movie)
	if err := c.BodyParser(response); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := s.movieService.UpdateMovie(movie); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update movie"})
	}
	response = mapper.MapMovieToMovieDTO(*movie)
	return c.JSON(movie)
}

// Удаление фильма
func (s *HttpServer) deleteMovie(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := s.movieService.DeleteMovie(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete movie"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
