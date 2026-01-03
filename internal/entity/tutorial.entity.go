package entity

import "github.com/google/uuid"

type Tutorial struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	Title        string    `json:"title" gorm:"not null"`
	Description  string    `json:"description" gorm:"not null"`
	TutorialType string    `json:"tutorial_type" gorm:"not null;check:tutorial_type IN ('fill-profile', 'create-pub')"`
	Duration     string    `json:"duration" gorm:"not null"`
	Image        string    `json:"image" gorm:"not null"`
	ButtonText   string    `json:"button_text" gorm:"not null"`
}
