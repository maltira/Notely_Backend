package dto

import "github.com/google/uuid"

type CategoryRequest struct {
	Name string `json:"name"`
}
type PublicationCategoryRequest struct {
	BackgroundColor string `json:"background_color"`
	TextColor       string `json:"text_color"`

	Category CategoryRequest
}

type CategoryUpdateRequest struct {
	ID   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
}
type PublicationCategoryUpdateRequest struct {
	ID            *uuid.UUID `json:"id"`
	PublicationID *uuid.UUID `json:"publication_id"`
	CategoryID    *uuid.UUID `json:"category_id"`

	BackgroundColor string `json:"background_color"`
	TextColor       string `json:"text_color"`
	DisplayOrder    int    `json:"display_order"`

	Category CategoryUpdateRequest
}
