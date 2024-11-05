package dto

type MovieDTO struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Year        int    `json:"year" validate:"required"`
}

//type UpdateMovieRequest struct {
//	Title       *string `json:"title"`
//	Description *string `json:"description"`
//	Year        *int    `json:"year"`
//}

//type MovieResponse struct {
//	ID          uint   `json:"id"`
//	Title       string `json:"title"`
//	Description string `json:"description"`
//	Year        int    `json:"year"`
//}
