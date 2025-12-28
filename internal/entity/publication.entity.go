package entity

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	Name string    `json:"name" gorm:"not null; unique"`
}

type Publication struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`

	BackgroundColor string `json:"background_color" gorm:"not null;size:7;default:'#F6F6F6'"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Связи
	User                  User                    `gorm:"foreignKey:UserID;onDelete:CASCADE"`
	PublicationCategories []PublicationCategories `gorm:"foreignKey:PublicationID;onDelete:CASCADE"`
}

type PublicationCategories struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey;not null"`
	PublicationID uuid.UUID `json:"publication_id" gorm:"type:uuid;not null;uniqueIndex:idx_pub_cat"`
	CategoryID    uuid.UUID `json:"category_id" gorm:"type:uuid;not null;uniqueIndex:idx_pub_cat"`

	BackgroundColor string `json:"background_color" gorm:"not null;size:7;default:'#F6F6F6'"`
	TextColor       string `json:"text_color" gorm:"not null;size:7;default:'#7D7D7D'"`
	DisplayOrder    int    `json:"display_order" gorm:"default:0"`

	// Связи
	Publication Publication `gorm:"foreignKey:PublicationID;onDelete:CASCADE"`
	Category    Category    `gorm:"foreignKey:CategoryID;onDelete:CASCADE"`
}
