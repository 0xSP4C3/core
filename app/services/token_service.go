package services

import (
	"context"

	"github.com/0xsp4c3/core/pkg/utils"
	"github.com/0xsp4c3/core/platform/cache"
	"github.com/0xsp4c3/core/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


func RenewTokens(id uuid.UUID) (int, string, error, *utils.Tokens) {

    // Create database connection.
    db, err := database.OpenDBConnection()
    if err != nil {
        // Return status 500 and database connection error.
        return fiber.StatusInternalServerError, "", err, nil
    }

    // Get user by ID.
    foundedUser, err := db.GetUserByID(id)
    if err != nil {
        // Return, if user not found.
        return fiber.StatusNotFound, "User with the given ID is not found.", err, nil
    }

    // Get role credentials from founded user.
    credentials, err := utils.GetCredentialsByRole(foundedUser.UserRole)
    if err != nil {
        // Return status 400 and error message.
        return fiber.StatusBadRequest, "", err, nil
    }

    // Generate JWT Access & Refresh tokens.
    tokens, err := utils.GenerateNewTokens(id.String(), credentials)
    if err != nil {
        // Return status 500 and token generation error.
        return fiber.StatusInternalServerError, "", err, nil
    }

    // Create a new Redis connection.
    connRedis, err := cache.RedisConnection()
    if err != nil {
        // Return status 500 and Redis connection error.
        return fiber.StatusInternalServerError, "", err, nil
    }

    // Save refresh token to Redis.
    errRedis := connRedis.Set(context.Background(), id.String(), tokens.Refresh, 0).Err()
    if errRedis != nil {
        // Return status 500 and Redis connection error.
        return fiber.StatusInternalServerError, "", errRedis, nil
    }

    return fiber.StatusOK, "", nil, tokens
}
