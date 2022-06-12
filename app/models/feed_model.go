package models

import (
	"time"

	"github.com/google/uuid"
)

type FeedQuery struct {
    StartAt     time.Time   `json:"start_at"`
    EndedAt     time.Time   `json:"ended_at"`
    CoinID      uuid.UUID   `json:"coin_id"`
    FeedRangeID uuid.UUID   `json:"feed_range_id"`
}

type Feed struct {
    ID          uuid.UUID   `db:"id" json:"id" validate:"required,uuid"`
    CreatedAt   time.Time   `db:"created_at" json:"created_at"`
    UpdatedAt   time.Time   `db:"upadted_at" json:"updated_at"`
    OpenBid     int64       `db:"open_bid" json:"open_bid" validate:"required"`
    CloseBid    int64       `db:"close_bid" json:"close_bid" validate:"required"`
    HighestBid  int64       `db:"highest_bid" json:"highest_bid" validate:"required"`
    LowestBid   int64       `db:"lowest_bid" json:"lowest_bid" validate:"required"`
    TotalTrade  int64       `db:"total_trade" json:"total_trade" validate:"required"`
    BaseVolume  int64       `db:"base_volume" json:"base_volume" validate:"required"`
    QuoteVolume int64       `db:"quote_volume" json:"qute_volume" validate:"required"`
    CoinID      uuid.UUID   `db:"coin_id" json:"coin_id" validate:"required,uuid"`
    FeedTimeID  uuid.UUID   `db:"feed_time_id" json:"feed_time_id" validate:"required,uuid"`
    FeedRangeID uuid.UUID   `db:"feed_range_id" json:"feed_range_id" validate:"required,uuid"`
}

type FeedTime struct {
    ID          uuid.UUID   `db:"id" json:"id" validate:"required,uuid"`
    CreatedAt   time.Time   `db:"created_at" json:"created_at"`
    UpdatedAt   time.Time   `db:"upadted_at" json:"updated_at"`
    StartAt     time.Time   `db:"start_at" json:"start_at" validate:"required"`
    EndedAt     time.Time   `db:"ended_at" json:"ended_at" validate:"required"`
}

type FeedRange struct {
    ID          uuid.UUID   `db:"id" json:"id" validate:"required,uuid"`
    CreatedAt   time.Time   `db:"created_at" json:"created_at"`
    UpdatedAt   time.Time   `db:"upadted_at" json:"updated_at"`
    Name        string      `db:"name" json:"name" validate:"required,lte=25"`
    Description string      `db:"description" json:"description" validate:"lte=2000"`
    IsEnabled   bool        `db:"is_enabled" json:"is_enabled"`
}
