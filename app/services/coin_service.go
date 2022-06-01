package services

import (
	"time"

	"github.com/0xsp4c3/core/app/models"
	"github.com/0xsp4c3/core/pkg/utils"
	"github.com/0xsp4c3/core/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


func GetCoins() (int, string, error, []models.Coin) {
    db, err := database.OpenDBConnection()
    if err != nil {
        return fiber.StatusInternalServerError, "", err, nil
    }

    coins, err := db.GetCoins()
    if err != nil {
        return fiber.StatusNotFound, "coin were not found.", err, nil
    }

    return fiber.StatusOK, "", nil, coins
}

func GetCoin(id uuid.UUID) (int, string, error, *models.Coin) {

    db, err := database.OpenDBConnection()
    if err != nil {
        return fiber.StatusInternalServerError, "", err, nil
    }

    coin, err := db.GetCoin(id)
    if err != nil {
        return fiber.StatusNotFound, "coin with the given ID is not found", err, nil
    }
    return fiber.StatusOK, "", nil, &coin
}

func GetCoinByExchangeID(id uuid.UUID) (int, string, error, []models.Coin) {
    db, err := database.OpenDBConnection()
    if err != nil {
        return fiber.StatusInternalServerError, "", err, nil
    }

    exchange, err := db.GetExchange(id)
    if err != nil {
        return fiber.StatusNotFound, "exchange with the given ID is not found", err, nil
    }

    coins, err := db.GetCoinsByExchangeID(exchange.ID)
    if err != nil {
        return fiber.StatusInternalServerError, "", err, nil
    }

    return fiber.StatusOK, "", nil, coins
}

func CreateCoin(c *models.Coin) (int, string, error) {
	// Create db connection
	db, err := database.OpenDBConnection()
	if err != nil {
        return fiber.StatusInternalServerError, "", err
	}

	// define current time
	time := time.Now()

	c.ID = uuid.New()
	c.CreatedAt = time
	c.UpdatedAt = time
	c.IsDeleted = false

	validate := utils.NewValidator()
	if err := validate.Struct(c); err != nil {
		// Return, if some fields are not valid.
        return fiber.StatusBadRequest, "", err
	}

	// Create coin
	if err := db.CreateCoin(c); err != nil {
        return fiber.StatusInternalServerError, "", err
	}

    return fiber.StatusCreated, "", nil
}

func UpdateCoin(c *models.Coin) (int, string, error) {
	db, err := database.OpenDBConnection()
	if err != nil {
        return fiber.StatusInternalServerError, "", err
	}

	foundedCoin, err := db.GetCoin(c.ID)
	if err != nil {
        return fiber.StatusNotFound, "Coin with this ID not found.", err
	}

	c.UpdatedAt = time.Now()

	validate := utils.NewValidator()
	if err := validate.Struct(c); err != nil {
        return fiber.StatusBadRequest, "", err
	}

	if err := db.UpdateCoin(foundedCoin.ID, c); err != nil {
        return fiber.StatusInternalServerError, "", err
	}

    return fiber.StatusNoContent, "", nil
}

func DeleteCoin(c *models.Coin) (int, string, error) {
    validate := utils.NewValidator()
    if err := validate.StructPartial(c); err != nil {
        return fiber.StatusBadRequest, "", err
    }

    db, err := database.OpenDBConnection()
    if err != nil {
        return fiber.StatusInternalServerError, "", err
    }

    foundedCoin, err := db.GetCoin(c.ID)
    if err != nil {
        return fiber.StatusNotFound, "Coin with this ID not found.", err
    }

    if err := db.DeleteCoin(foundedCoin.ID); err != nil {
        return fiber.StatusInternalServerError, "", err
    }

    return fiber.StatusNoContent, "", nil
}

