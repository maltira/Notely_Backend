package dto

import (
	"github.com/google/uuid"
)

type CategoryRequest struct {
	Name string `json:"name"`
}
type PublicationCategoryRequest struct {
	BackgroundColor string `json:"background_color"`
	TextColor       string `json:"text_color"`

	Category CategoryRequest `json:"category"`
}

type PublicationRequest struct {
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	UserID          uuid.UUID `json:"user_id"`
	BackgroundColor string    `json:"background_color"`

	Categories []PublicationCategoryRequest `json:"categories"`
}

type PublicationUpdateRequest struct {
	ID              uuid.UUID                     `json:"id"`
	Title           *string                       `json:"title"`
	Description     *string                       `json:"description"`
	BackgroundColor *string                       `json:"background_color"`
	Categories      *[]PublicationCategoryRequest `json:"categories"`
}
