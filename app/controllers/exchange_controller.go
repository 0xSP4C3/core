package controllers

import (
	"time"

	"github.com/0xsp4c3/core/app/models"
	"github.com/0xsp4c3/core/pkg/repository"
	"github.com/0xsp4c3/core/pkg/utils"
	"github.com/0xsp4c3/core/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

// GetExchange func gets exchange by given ID or 404 error.
// @Description Get exchange by given ID.
// @Summary get exchange by given ID
// @Tags Exchange
// @Accept json
// @Produce json
// @Param id path string true "Exchange ID"
// @Success 200 {object} models.Exchange
// @Router /v1/exchange/{id} [get]
func GetExchange(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }

    db, err := database.OpenDBConnection()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }

    exchange, err := db.GetExchange(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error":    true,
            "msg":      "exchange with the given ID is not found",
            "exchange": nil,
        })
    }

    return c.JSON(fiber.Map{
        "error":    false,
        "msg":      nil,
        "exchange": exchange,
    })
}

func CreateExchange(c *fiber.Ctx) error {
    currentTime := time.Now().Unix()

    claims, err := utils.ExtractTokenMetadata(c)
    if err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }

    expireTime := claims.Expires

    if currentTime > expireTime {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error":    true,
            "msg":      "Unauthorized. Token expired.",
        })
    }

    credential := claims.Credentials[repository.ExchangeCreateCredential]
    if !credential {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error":    true,
            "msg":      "Forbidden. Permission denied.",
        })
    }
    
    exchange := &models.Exchange{}

    if err := c.BodyParser(exchange); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }

    db, err := database.OpenDBConnection()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }

    time := time.Now()

    validate := utils.NewValidator()

    exchange.ID = uuid.New()
    exchange.CreatedAt = time
    exchange.UpdatedAt = time
    exchange.IsBlocked = false
    exchange.IsEnabled = true

    if err := validate.Struct(exchange); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":    true,
            "msg":      utils.ValidatorErrors(err),
        })
    }

    if err := db.CreateExchange(exchange); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "error":    false,
        "msg":      "Exchange created!",
        "exchange": exchange,
    })
}
