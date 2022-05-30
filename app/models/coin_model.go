package models

import (
	"time"

	"github.com/google/uuid"
)


type Coin struct {
    ID          uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
    CreatedAt   time.Time `db:"created_at" json:"created_at"`
    UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
    ExchangeId  uuid.UUID `db:"exchange_id" json:"exchange_id" validate:"required,uuid"`
    Name        string `db:"name" json:"name" validate:"required,gte=2,lte=25"`
    Code        string `db:"code" json:"code" validate:"required,gte=2,lte=10"`
    Description string `db:"description" json:"description" validate:"lte=255"`
    IsDeleted   bool `db:"is_deleted" json:"is_deleted" validate:"required"`
    ImageUri    string `db:"image_uri" json:"image_uri"`
    CoinUri     CoinUri `db:"coin_uri" json:"coin_uri" validate:"required,dive"`
}

type CoinUri struct {
    CoinID      uuid.UUID `db:"coin_id" json:"coin_id" validate:"required,uuid"`
    Uri         string `db:"uri" json:"uri" validate:"required"`
    CreatedAt   time.Time `db:"created_at" json:"created_at"`
    UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
