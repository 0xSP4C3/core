package controllers

import (
	"github.com/0xsp4c3/core/app/models"
	"github.com/gofiber/fiber/v2"
)


func QueryFeeds(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "error": false,
        "msg":   nil,
        "feeds": nil, // Array results.
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
func GetFeedRange(c *fiber.Ctx) error {
    feedRanges := []models.FeedRange{}

    return c.JSON(fiber.Map{
        "error": false,
        "msg":   nil,
        "count": len(feedRanges),
        "feed_ranges": nil,
    })
}

func CreateFeed(c *fiber.Ctx) error {
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "error": false,
        "msg":   "Feed Created.",
    })
}

func UpdateFeed(c *fiber.Ctx) error {
    return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
        "error": false,
        "msg":   nil,
    })
}

func DeleteFeed(c *fiber.Ctx) error {
    return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
        "error": false,
        "msg":   nil,
    })
}


