package controllers

import (
	"github.com/0xsp4c3/core/platform/database"
	"github.com/gofiber/fiber/v2"
)

// GetExchange funcs gets all exists exchanges.
// @Description Get all exists exchanges.
// @Summary get all exists exchanges
// @Tags exchange
// @Accept json
// @Produce json
// @Success 200 {array} models.Exchange
// @Router /v1/exchange [get]
func GetExchanges(c *fiber.Ctx) error {
    db, err := database.OpenDBConnection()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }
    exchanges, err := db.GetExchanges()
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error":    true,
            "msg":      "exchanges were not found.",
        })
    }

    return c.JSON(fiber.Map{
        "error":    false,
        "msg":      nil,
        "count":    len(exchanges),
        "exchanges":exchanges,
    })
}
