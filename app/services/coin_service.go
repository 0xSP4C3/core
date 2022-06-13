package services

import (
	"time"

	"github.com/0xsp4c3/core/app/models"
	"github.com/0xsp4c3/core/pkg/utils"
	"github.com/0xsp4c3/core/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CoinService struct {}


func (s *CoinService)GetCoins() (statusCode int, message string, err error, results []models.Coin) {
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

func (s *CoinService)GetCoin(id uuid.UUID) (statusCode int, message string, err error, result *models.Coin) {

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

func (s *CoinService)GetCoinByExchangeID(id uuid.UUID) (statusCode int, message string, err error, results []models.Coin) {
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

func (s *CoinService)CreateCoin(c *models.Coin) (statusCode int, message string, err error) {
	// Create db connection
	db, err := database.OpenDBConnection()
	if err != nil {
        return fiber.StatusInternalServerError, "", err
	}

	// define current time
	currentTime := time.Now()

	c.ID = uuid.New()
	c.CreatedAt = currentTime
	c.UpdatedAt = currentTime
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

func (s *CoinService)UpdateCoin(c *models.Coin) (statusCode int, message string, err error) {
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

func (s *CoinService)DeleteCoin(c *models.Coin) (statusCode int, message string, err error) {
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

