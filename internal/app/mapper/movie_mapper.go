package mapper

import (
	"github.com/Ablyamitov/cinema/internal/app/db/models"
	"github.com/Ablyamitov/cinema/internal/app/dto"
)

func MapMovieDTOToMovie(request dto.MovieDTO) models.Movie {
	return models.Movie{
		Title:       request.Title,
		Description: request.Description,
		Year:        request.Year,
	}
}

func MapMovieToMovieDTO(movie models.Movie) dto.MovieDTO {
	return dto.MovieDTO{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Year:        movie.Year,
	}
}

func MapMoviesToMoviesDto(movies []models.Movie) []dto.MovieDTO {
	moviesDTO := make([]dto.MovieDTO, len(movies))
	for i, movie := range movies {
		moviesDTO[i] = MapMovieToMovieDTO(movie)
	}
	return moviesDTO
}
