package service

import (
	"github.com/Ablyamitov/cinema/internal/app/db/models"
)

type MovieService interface {
	CreateMovie(movie *models.Movie) error
	GetMovieByID(id uint) (*models.Movie, error)
	UpdateMovie(movie *models.Movie) error
	DeleteMovie(id uint) error
	GetAllMovies() ([]models.Movie, error)
}

var _ MovieService = (*movieService)(nil)

type movieService struct {
}

func NewMovieService() MovieService {
	return &movieService{}
}

func (s *movieService) CreateMovie(movie *models.Movie) error {
	if err := movie.Create(); err != nil {
		return err
	}
	return nil
}

func (s *movieService) GetMovieByID(id uint) (*models.Movie, error) {
	movie := models.Movie{ID: id}
	if err := movie.Find(); err != nil {
		return nil, err
	}
	return &movie, nil
}

func (s *movieService) UpdateMovie(movie *models.Movie) error {
	if err := movie.Save(); err != nil {
		return err
	}
	return nil
}

func (s *movieService) DeleteMovie(id uint) error {
	movie, err := s.GetMovieByID(id)
	if err != nil {
		return err
	}
	if err := movie.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *movieService) GetAllMovies() ([]models.Movie, error) {
	return make([]models.Movie, 0), nil
}
