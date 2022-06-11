package controllers

import "github.com/gofiber/fiber/v2"


func QueryFeeds(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "error": false,
        "msg":   nil,
        "feeds": nil, // Array results.
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


