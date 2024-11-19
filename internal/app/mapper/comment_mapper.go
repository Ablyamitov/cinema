package mapper

import (
	"github.com/Ablyamitov/cinema/internal/app/db/models"
	"github.com/Ablyamitov/cinema/internal/app/dto"
)

func MapCommentDTOToComment(request dto.CommentDTO) models.Comment {
	return models.Comment{
		Author: request.Author,
		Title:  request.Title,
	}
}

func MapCommentToCommentDTO(comment models.Comment) dto.CommentDTO {
	return dto.CommentDTO{
		ID:        comment.ID,
		Author:    comment.Author,
		Title:     comment.Title,
		CreatedAt: comment.CreatedAt,
	}
}

func MapCommentsToCommentsDto(comments []models.Comment) []dto.CommentDTO {
	commentsDTO := make([]dto.CommentDTO, len(comments))
	for i, comment := range comments {
		commentsDTO[i] = MapCommentToCommentDTO(comment)
	}
	return commentsDTO
}
