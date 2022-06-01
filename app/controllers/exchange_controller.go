package controllers

import (
	"time"

	"github.com/0xsp4c3/core/app/models"
	"github.com/0xsp4c3/core/app/services"
	"github.com/0xsp4c3/core/pkg/repository"
	"github.com/0xsp4c3/core/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetExchange funcs gets all exists exchanges.
// @Description Get all exists exchanges.
// @Summary get all exists exchanges
// @Tags Exchanges
// @Accept json
// @Produce json
// @Success 200 {array} models.Exchange
// @Router /v1/exchanges [get]
func GetExchanges(c *fiber.Ctx) error {
    statusCode, message, err, exchanges := services.GetExchanges()
    if err != nil {
        var msg string
        if message == "" {
            msg = err.Error()
        } else {
            msg = message
        }
        return c.Status(statusCode).JSON(fiber.Map{
            "error":    true,
            "msg":      msg,
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
    statusCode, message, err, exchange := services.GetExchange(id)
    if err != nil {
        var msg string
        if message == "" {
            msg = err.Error()
        } else {
            msg = message
        }
        return c.Status(statusCode).JSON(fiber.Map{
            "error":    true,
            "msg":      msg,
            "exchange": nil,
        })
    }

    return c.JSON(fiber.Map{
        "error":    false,
        "msg":      nil,
        "exchange": exchange,
    })
}

// CreateExchange func for creates a new exchange.
// @Description Create a new exchange.
// @Summary create a new exchange
// @Tags Exchange
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param description body string true "Description"
// @Success 200 {object} models.Exchange
// @Security ApiKeyAuth
// @Router /v1/exchange [post]
func CreateExchange(c *fiber.Ctx) error {
    currentTime := time.Now().Unix()

    claims, err := utils.ExtractTokenMetadata(c)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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

    statusCode, message, err := services.CreateExchange(exchange)
    if err != nil {
        var msg string
        if message == "" {
            msg = err.Error()
        } else {
            msg = message
        }
        return c.Status(statusCode).JSON(fiber.Map{
            "error":    true,
            "msg":      msg,
        })
    }
    return c.Status(statusCode).JSON(fiber.Map{
        "error":    false,
        "msg":      message,
        "exchange": exchange,
    })
}

// UpdateExchange func for updates exchange by given ID.
// @Description Update exchange.
// @Summary update exchange
// @Tags Exchange
// @Accept json
// @Produce json
// @Param id body string true "Exchange ID"
// @Param name body string true "Name"
// @Param description body string true "Description"
// @Param uri body string true "Uri"
// @Param is_enabled body boolean true "Is Enabled"
// @Param is_blocked body boolean true "Is Blocked"
// @Success 202 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/exchange [put]
func UpdateExchange(c *fiber.Ctx) error {
    currentTime := time.Now().Unix()

    claims, err := utils.ExtractTokenMetadata(c)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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
    
    credential := claims.Credentials[repository.ExchangeUpdateCredential]
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

    statusCode, message, err := services.UpdateExchange(exchange)
    if err != nil {
        var msg string
        if message == "" {
            msg = err.Error()
        } else {
            msg = message
        }
        return c.Status(statusCode).JSON(fiber.Map{
            "error":    true,
            "msg":      msg,
        })
    }
    return c.SendStatus(statusCode)
}

// DeleteExchange func for deletes exchange by given ID.
// @Description Delete Exchange by given ID.
// @Summary delete exchange by given ID
// @Tags Exchange
// @Accept json
// @Produce json
// @Param id body string true "Exchange ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/exchange [delete]
func DeleteExchange(c *fiber.Ctx) error {
    currentTime := time.Now().Unix()

    claims, err := utils.ExtractTokenMetadata(c)
    if err != nil{
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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
    
    credential := claims.Credentials[repository.ExchangeDeleteCredential]
    if !credential {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error":    true,
            "msg":      "Forbidden. Permission denied.",
        })
    }

    exchange := &models.Exchange{}

    if err := c.BodyParser(exchange); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }

    statusCode, message, err := services.DeleteExchange(exchange)
    if err != nil {
        var msg string
        if message == "" {
            msg = err.Error()
        } else {
            msg = message
        }
        return c.Status(statusCode).JSON(fiber.Map{
            "error":    true,
            "msg":      msg,
        })
    }
    return c.SendStatus(statusCode)
}
