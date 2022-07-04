package services

import (
	"github.com/0xsp4c3/core/app/models"
	"github.com/0xsp4c3/core/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FeedService struct {}

// WARN: Don't use this.
func (s *FeedService) GetFeeds() (statusCode int, message string, err error, results []models.Feed) {
    db, err := database.OpenDBConnection()
    if err != nil {
        return fiber.StatusInternalServerError, "", err, nil
    }

    feeds, err := db.GetFeeds()
    if err != nil {
        return fiber.StatusNotFound, "feeds were not found.", err, nil 
    }

    return fiber.StatusOK, "", nil, feeds
}

func (s *FeedService) QueryFeedsByTimeRange(range_id uuid.UUID) (statusCode int, message string, err error, results []models.Feed) {
    db, err := database.OpenDBConnection()
    if err != nil {
        return fiber.StatusInternalServerError, "", err, nil
    }

    feed_range, err := db.GetFeedRange(range_id)
    if err != nil {
        return fiber.StatusNotFound, "Feed range not found.", err, nil
    }
    
    feeds, err := db.QueryFeedsByTimeRange(feed_range.ID)
    if err != nil {
        return fiber.StatusInternalServerError, "", err, nil
    }

    return fiber.StatusOK, "", nil, feeds
}

func (s *FeedService) QueryFeedsByCoinId(coin_id uuid.UUID) (statusCode int, message string, err error, results []models.Feed) {
    db, err := database.OpenDBConnection()
    if err != nil {
        return fiber.StatusInternalServerError, "", err, nil
    }

    coin, err := db.GetCoin(coin_id)
    if err != nil {
        return fiber.StatusNotFound, "Coin not found.", err, nil
    }

    feeds, err := db.QueryFeedsByCoinId(coin.ID)
    if err != nil {
        return fiber.StatusInternalServerError, "", err, nil
    }

    return fiber.StatusOK, "", err, feeds
}

func (s *FeedService) QueryFeedsByCoinAndRangeId(range_id, coin_id uuid.UUID) (statusCode int, message string, err error, results []models.Feed) {
    db, err := database.OpenDBConnection()
    if err != nil {
        return fiber.StatusInternalServerError, "", err, nil
    }
    feed_range, err := db.GetFeedRange(range_id)
    if err != nil {
        return fiber.StatusNotFound, "Feed Range not found.", err, nil
    }

    coin, err := db.GetCoin(coin_id)
    if err != nil {
        return fiber.StatusNotFound, "Coin not found.", err, nil
    }

    feeds, err := db.QueryFeedsByCoinAndRangeId(feed_range.ID, coin.ID)
    if err != nil {
        return fiber.StatusNotFound, "Feeds not found.", err, nil
    }

    return fiber.StatusOK, "", nil, feeds
}
