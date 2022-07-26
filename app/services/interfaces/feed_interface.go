package interfaces

import "github.com/0xsp4c3/core/app/models"

type FeedService interface {
    GetFeedRanges() (statusCode int, message, string, err error, results []models.TimeFrame)
}

