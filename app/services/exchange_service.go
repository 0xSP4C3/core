package services

import (
	"time"

	"github.com/0xsp4c3/core/app/models"
	"github.com/0xsp4c3/core/pkg/utils"
	"github.com/0xsp4c3/core/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetExchanges() (int, string, error, []models.Exchange) {
    db, err := database.OpenDBConnection()
    if err != nil {
        return fiber.StatusInternalServerError, "", err, nil
    }

    exchanges, err := db.GetExchanges()
    if err != nil {
        return fiber.StatusNotFound, 
            "exchanges were not found.",
            err, 
            nil
    }

    return fiber.StatusOK, "", nil, exchanges
}

func GetExchange(id uuid.UUID) (int, string, error, *models.Exchange) {
    db, err := database.OpenDBConnection()
    if err != nil {
        return fiber.StatusInternalServerError, "", err, nil
    }

    exchange, err := db.GetExchange(id)
    if err != nil {
        return fiber.StatusNotFound, 
            "exchange with given ID is not found.", 
            err, 
            nil
    }

    return fiber.StatusOK, "", nil, &exchange
}

func CreateExchange(e *models.Exchange) (int, string, error) {
    db, err := database.OpenDBConnection()
    if err != nil {
        return fiber.StatusInternalServerError, "", err
    }

    time := time.Now()

    e.ID = uuid.New()   
    e.CreatedAt = time
    e.UpdatedAt = time
    e.IsBlocked = false
    e.IsEnabled = true

    validate := utils.NewValidator()
    if err := validate.Struct(e); err != nil {
        return fiber.StatusBadRequest, "", err
    }

    if err := db.CreateExchange(e); err != nil {
        return fiber.StatusInternalServerError, "", err
    }

    return fiber.StatusOK, "Exchange created!", nil

}

func UpdateExchange(e *models.Exchange) (int, string, error) {
    db, err := database.OpenDBConnection()
    if err != nil {
        return fiber.StatusInternalServerError, "", err
    }

    item, err := db.GetExchange(e.ID)
    if err != nil {
        return fiber.StatusNotFound, 
            "exhcange with this ID not found",
            err
    }

    e.UpdatedAt = time.Now()

    validate := utils.NewValidator()
    if err := validate.StructPartial(e); err != nil {
        return fiber.StatusBadRequest,
            "",
            err
    }

    if err := db.UpdateExchange(item.ID, e); err != nil {
        return fiber.StatusInternalServerError, "", err
    }

    return fiber.StatusNoContent, "", nil
}

func DeleteExchange(e *models.Exchange) (int, string, error) {

    validate := utils.NewValidator()
    if err := validate.StructPartial(e, "id"); err != nil {
        return fiber.StatusBadRequest, "", err
    }

    db, err := database.OpenDBConnection()
    if err != nil {
        return fiber.StatusInternalServerError, "", err
    }

    item, err := db.GetExchange(e.ID)
    if err != nil {
        return fiber.StatusNotFound, 
            "Not found. Exchange with this ID not found.", 
            err
    }

    if err := db.DeleteExchange(item.ID); err != nil {
        return fiber.StatusInternalServerError,
            "",
            err
    }
    return fiber.StatusNoContent, "", nil
}
