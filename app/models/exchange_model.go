package models

import (
	"time"

	"github.com/google/uuid"
)



type Exchange struct {
    ID          uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
    CreatedAt   time.Time `db:"created_at" json:"created_at"`
    UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
    Name        string    `db:"name" json:"name" validate:"required,lte=25"`
    Description string    `db:"description" json:"description" validate:"lte=255"`
    Uri         string    `db:"uri" json:"uri"`
    // Enable Crawling?
    IsEnabled   bool      `db:"is_enabled" json:"is_enabled" validate:"required"`
    // We got blocked?
    IsBlocked   bool      `db:"is_blocked" json:"is_blocked" validate:"required"`
}
