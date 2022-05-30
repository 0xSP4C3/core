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

// CreateCoin func for creates a new coin.
// @Description Create a new coin.
// @Summary create a new coin
// @Tags coin
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param code body string true "Code"
// @Param exchange_id body string true "Exchange ID"
// @Param coin_uri body models.CoinUri true "Coin Uri"
// @Success 200 {object} models.Coin
// @Security ApiKeyAuth
// @Router /v1/coin [post]
func CreateCoin(c *fiber.Ctx) error {
    // Get current time
    currentTime := time.Now().Unix()
    
    // Check Claims
    claims, err := utils.ExtractTokenMetadata(c)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":  err.Error(),
        })
    }
    
    // Get token expires time from JWT token.
    expireTime := claims.Expires

    // Check, if token is expired.
    if currentTime > expireTime {
        //Return status 401 and unauthorized error message
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": true,
            "msg": "Unauthorized. Token expired.",
        })
    }

    // Permission check
    // Get permission list from JWT token.
    credentials := claims.Credentials[repository.CoinCreateCredential]

    // Forbidden request doesn't have enough premission.
    if !credentials {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error": true,
            "msg": "Forbidden. Permission denied.",
        })
    }

    // create new coin struct
    coin := &models.Coin{}

    // Check, if received data is valid.
    if err := c.BodyParser(coin); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": true,
            "msg":  err.Error(),
        })
    }

    // Create db connection
    db, err := database.OpenDBConnection()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":  err.Error(),
        })
    }

    // define current time
    time := time.Now()

    validate := utils.NewValidator()

    coin.ID = uuid.New()
    coin.CreatedAt = time
    coin.UpdatedAt = time
    coin.IsDeleted = false

    // Validate coin fields.
    if err := validate.Struct(coin); err != nil {
        // Return, if some fields are not valid.
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": true,
            "msg": utils.ValidatorErrors(err),
        })
    }

    // Create coin
    if err := db.CreateCoin(coin); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":  err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "error":    false,
        "msg":      "Coin Created!",
        "coin":     coin,
    })
}
