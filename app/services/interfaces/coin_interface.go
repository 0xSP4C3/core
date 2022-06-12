package interfaces

import (
	"github.com/0xsp4c3/core/app/models"
    "github.com/google/uuid"
)

type CoinService interface {
   GetCoins() (statusCode int, msg string, err error, result []models.Coin)
   GetCoin(id uuid.UUID) (statusCode int, message string, err error, result *models.Coin)
   GetCoinByExchangeID(id uuid.UUID) (statusCode int, message string, err error, result []models.Coin)
   CreateCoin(c *models.Coin) (statusCode int, message string, err error)
   DeleteCoin(c *models.Coin) (statusCode int, message string, err error)
}
