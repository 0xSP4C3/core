package models

import (
	"time"

	"github.com/google/uuid"
)

type Feed struct {
    ID          uuid.UUID
    CreatedAt   time.Time
    UpdatedAt   time.Time
    OpenBid     int64
    CloseBid    int64
    HighestBid  int64
    LowestBid   int64
    TotalTrade  int64
    BaseVolume  int64
    QuoteVolume int64
    CoinID      uuid.UUID
    FeedTimeID  uuid.UUID
    FeedRangeID uuid.UUID
}

type FeedTime struct {
    ID          uuid.UUID
    CreatedAt   time.Time
    UpdatedAt   time.Time
    StartAt     time.Time
    EndedAt     time.Time
}

type FeedRange struct {
    ID          uuid.UUID
    CreatedAt   time.Time
    UpdatedAt   time.Time
    Name        string
    Description string
    IsEnabled   bool
}
