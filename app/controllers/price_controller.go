package controllers

import (
	"time"

	"github.com/0xsp4c3/core/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)



func GetPriceByCoinID(c *fiber.Ctx) error {

    id, err := uuid.Parse(c.Params("id")) 
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }

    validate, err := utils.ExtractTokenMetadata(c)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }

    currentTime := time.Now().Unix()

    expireTime := validate.Expires

    if currentTime > expireTime {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": true,
            "msg":   "Unauthorized. Token expired.",
        })
    }

}
