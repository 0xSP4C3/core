package controllers

import (
	"github.com/0xsp4c3/core/app/models"
	"github.com/gofiber/fiber/v2"
)

// QueryFeeds method to get feeds.
// @Description Query selected Feeds
// @Summary Get feed's all available time ranges.
// @Tags Feed
// @Accept json
// @Produce json
// @Success 200 {array} models.Feed
// @Router /v1/queryfeeds [get]
func QueryFeeds(c *fiber.Ctx) error {
    feeds := []models.Feed{}
    return c.JSON(fiber.Map{
        "error": false,
        "msg":   nil,
        "feeds": feeds, // Array results.
    })
}

// FeedRange method to get feed available range.
// @Description Get Feed Ranges.
// @Summary Get feed's all available time ranges.
// @Tags Feed
// @Accept json
// @Produce json
// @Success 200 {array} models.FeedTimeRange
// @Router /v1/feedranges [get]
func GetFeedRanges(c *fiber.Ctx) error {
    timeFrame := []models.TimeFrame{}

    return c.JSON(fiber.Map{
        "error": false,
        "msg":   nil,
        "count": len(timeFrame),
        "feed_ranges": nil,
    })
}

// CreateFeed method to create Feed
// @Description Create Feed
// @Summary Create Feed.
// @Tags Feed
// @Accept json
// @Produce json
// @Success 204
// @Router /v1/feedranges [get]
func CreateFeed(c *fiber.Ctx) error {
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "error": false,
        "msg":   "Feed Created.",
    })
}

// UpdateFeed method to update feed
// @Description Update Feed
// @Summary Update Feed.
// @Tags Feed
// @Accept json
// @Produce json
// @Success 204
// @Router /v1/updatefeed [get]
func UpdateFeed(c *fiber.Ctx) error {
    return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
        "error": false,
        "msg":   nil,
    })
}

func GetTimeFrame(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "error": false,
        "msg":   nil,
        "time_frame": nil,
    })
}

func CreateTimeFrame(c *fiber.Ctx) error {
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "error": false,
        "msg":   "TimeFrame Created.",
    })
}

func UpdateTimeFrame(c *fiber.Ctx) error {
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "error": false,
        "msg":   "TimeFrame Updated.",
    })
}


