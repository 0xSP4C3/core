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

// GetCoins funcs gets all exists coins.
// @Description Get all exists coins.
// @Summary get all exists coins
// @Tags Coins
// @Accept json
// @Produce json
// @Success 200 {array} models.Coin
// @Router /v1/coins [get]
func GetCoins(c *fiber.Ctx) error {
    db ,err := database.OpenDBConnection()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }

    coins, err := db.GetCoins()
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error":    true,
            "msg":      "coins were not found",
        })
    }

    return c.JSON(fiber.Map{
        "error":    false,
        "msg":      nil,
        "count":    len(coins),
        "coins":    coins,
    })
}

// GetCoin func gets coin by given ID or 404 error.
// @Description Get coin by given ID.
// @Summary get coin by given ID
// @Tags Coin
// @Accept json
// @Produce json
// @Param id path string true "Coin ID"
// @Success 200 {object} models.Coin
// @Router /v1/coin/{id} [get]
func GetCoin(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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

    coin, err := db.GetCoin(id)
    if err != nil {
      return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
        "error":    true,
        "msg":      "coin with the given ID is not found",
        "coin":     nil,
      })
    }

    return c.JSON(fiber.Map{
      "error":  false,
      "msg":    nil,
      "coin":   coin,
    })
}

// GetCoinByExchangeId func gets coin by given ID of Exchange or 404 error.
// @Description Get coin by given exchange ID
// @Summary get coin by given exhcange ID
// @Tags Coin
// @Accept json
// @Produce json
// @Param id path string true "Exchange ID"
// @Success 200 {object} models.Coin
// @route /v1/coinbyexchangeid/{id} [get]
func GetCoinByExchangeID(c *fiber.Ctx) error {
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

    coin, err := db.GetCoinByExchangeID(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error":    true,
            "msg":      "coin with the given exchange ID is not found",
            "coin":     nil,
        })
    }

    return c.JSON(fiber.Map{
        "error":    false,
        "msg":      nil,
        "coin":     coin,
    })
}

// CreateCoin func for creates a new coin.
// @Description Create a new coin.
// @Summary create a new coin
// @Tags Coin
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
            "error":    true,
            "msg":      err.Error(),
        })
    }
    
    // Get token expires time from JWT token.
    expireTime := claims.Expires

    // Check, if token is expired.
    if currentTime > expireTime {
        //Return status 401 and unauthorized error message
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error":    true,
            "msg":      "Unauthorized. Token expired.",
        })
    }

    // Permission check
    // Get permission list from JWT token.
    credentials := claims.Credentials[repository.CoinCreateCredential]

    // Forbidden request doesn't have enough premission.
    if !credentials {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error":    true,
            "msg":      "Forbidden. Permission denied.",
        })
    }

    // create new coin struct
    coin := &models.Coin{}

    // Check, if received data is valid.
    if err := c.BodyParser(coin); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }

    // Create db connection
    db, err := database.OpenDBConnection()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
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
            "error":    true,
            "msg":      utils.ValidatorErrors(err),
        })
    }

    // Create coin
    if err := db.CreateCoin(coin); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "error":    false,
        "msg":      "Coin Created!",
        "coin":     coin,
    })
}

// UpdateCoin func for updates coin by given ID.
// @Description Update coin.
// @Summary update coin
// @Tags Coin
// @Accept json
// @Produce json
// @Param id body string true "Coin ID"
// @Param name body string true "Name"
// @Param code body string true "Code"
// @Param description body string true "Description"
// @Param exchange_id body string true "Exchange ID"
// @Param coin_uri body models.CoinUri true "Coin Uri"
// @Success 202 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/coin [put]
func UpdateCoin(c *fiber.Ctx) error {
    currentTime := time.Now().Unix()

    claims, err:= utils.ExtractTokenMetadata(c)
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

    credential := claims.Credentials[repository.CoinUpdateCredential]

    if !credential {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error":    true,
            "msg":      "Forbidden. Permission denied.",
        })
    }

    coin := &models.Coin{}

    if err := c.BodyParser(coin); err != nil {
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

    foundedCoin, err := db.GetCoin(coin.ID)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error":    true,
            "msg":      "coin with this ID not found",
        })
    }
    
    coin.UpdatedAt = time.Now()
    
    validate := utils.NewValidator()

    if err := validate.Struct(coin); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }

    if err := db.UpdateCoin(foundedCoin.ID, coin); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "error":    false,
        "msg":      nil,
    })
}

// DeleteCoin func for deletes coin by given ID.
// @Description Delete coin by given ID.
// @Summary delete coin by given ID
// @Tags Coin
// @Accept json
// @Produce json
// @Param id body string true "Coin ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/coin [delete]
func DeleteCoin(c *fiber.Ctx) error {
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

    credential := claims.Credentials[repository.CoinDeleteCredential]

    if !credential {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error":    true,
            "msg":      "Forbidden. Permission denied.", 
        })
    }

    coin := &models.Coin{}
    
    if err := c.BodyParser(coin); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }
    
    validate := utils.NewValidator()

    if err := validate.StructPartial(coin, "id"); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":    true,
            "msg":      utils.ValidatorErrors(err),
        })
    }

    db, err := database.OpenDBConnection()

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }

    foundedCoin, err := db.GetCoin(coin.ID)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error":    true,
            "msg":      "Not found. Coin with this ID not found.",
        })
    }

    if err:= db.DeleteBook(foundedCoin.ID); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error":    true,
            "msg":      err.Error(),
        })
    }

    return c.SendStatus(fiber.StatusNoContent)
}
