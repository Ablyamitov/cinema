package service

import (
	"github.com/Ablyamitov/cinema/internal/app/db/models"
	"gorm.io/gorm"
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
	db *gorm.DB
}

func NewMovieService(db *gorm.DB) MovieService {
	return &movieService{db: db}
}

func (s *movieService) CreateMovie(movie *models.Movie) error {
	return s.db.Create(movie).Error
}

func (s *movieService) GetMovieByID(id uint) (*models.Movie, error) {
	var movie models.Movie
	if err := s.db.First(&movie, id).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (s *movieService) UpdateMovie(movie *models.Movie) error {
	return s.db.Save(movie).Error
}

func (s *movieService) DeleteMovie(id uint) error {
	return s.db.Delete(&models.Movie{}, id).Error
}

func (s *movieService) GetAllMovies() ([]models.Movie, error) {
	var movies []models.Movie
	if err := s.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}
