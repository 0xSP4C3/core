package interfaces

import (
	"github.com/0xsp4c3/core/app/models"
	"github.com/google/uuid"
)


type ExchangeInterface interface {
    GetExchanges() (statusCode int, msg string, err error, result []models.Exchange)
    GetExchange(id uuid.UUID) (statusCode int, msg string, err error, resut *models.Exchange)
    CreateExchange(e *models.Exchange) (statusCode int, msg string, err error)
    UpdateExchange(e *models.Exchange) (statusCode int, msg string, err error)
    DeleteExchange(e *models.Exchange) (statusCode int, msg string, err error)
}
