package entity

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;not null;uniqueIndex:idx_subscription_pair"`
	TargetID uuid.UUID `json:"target_id" gorm:"type:uuid;not null;uniqueIndex:idx_subscription_pair"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Связи
	SubscriberUser *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	TargetUser     *User `gorm:"foreignKey:TargetID;constraint:OnDelete:CASCADE"`
}
