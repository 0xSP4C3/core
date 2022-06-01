package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/0xsp4c3/core/app/models"
	"github.com/0xsp4c3/core/pkg/utils"
	"github.com/0xsp4c3/core/platform/cache"
	"github.com/0xsp4c3/core/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserSignUp(signUp *models.SignUp) (statusCode int, message string, err error, result *models.User) {
	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(signUp); err != nil {
        msg, jsonErr := json.Marshal(utils.ValidatorErrors(err))
        if jsonErr != nil {
            return fiber.StatusInternalServerError, "", err, nil
        }
		// Return, if some fields are not valid.
        return fiber.StatusBadRequest, string(msg), err, nil
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
        return fiber.StatusInternalServerError, "", err, nil
	}

	// Checking role from sign up data.
	role, err := utils.VerifyRole(signUp.UserRole)
	if err != nil {
		// Return status 400 and error message.
        return fiber.StatusBadRequest, "", err, nil
	}

	// Create a new user struct.
	user := &models.User{}

	// Set initialized default data for user:
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.Email = signUp.Email
	user.PasswordHash = utils.GeneratePassword(signUp.Password)
	user.UserStatus = 1 // 0 == blocked, 1 == active
	user.UserRole = role

	// Validate user fields.
	if err := validate.Struct(user); err != nil {
        msg, jsonErr := json.Marshal(utils.ValidatorErrors(err))
        if jsonErr != nil {
            return fiber.StatusInternalServerError, "", err, nil
        }
		// Return, if some fields are not valid.
        return fiber.StatusBadRequest, string(msg), err, nil
	}

	// Create a new user with validated data.
	if err := db.CreateUser(user); err != nil {
		// Return status 500 and create user process error.
        return fiber.StatusInternalServerError, "", err, nil
	}

	// Delete password hash field from JSON view.
	user.PasswordHash = ""
    return fiber.StatusOK, "", nil, user
}

func UserSignIn(signIn *models.SignIn) (statusCode int, message string, err error, result *utils.Tokens) {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
        return fiber.StatusInternalServerError, "", err, nil
	}

	// Get user by email.
	foundedUser, err := db.GetUserByEmail(signIn.Email)
	if err != nil {
		// Return, if user not found.
        return fiber.StatusNotFound, "User with the given email is not found.", err, nil
	}

	// Compare given user password with stored in found user.
	compareUserPassword := utils.ComparePasswords(foundedUser.PasswordHash, signIn.Password)
	if !compareUserPassword {
		// Return, if password is not compare to stored in database.
        return fiber.StatusBadGateway, "Wrong user email address or password.", err, nil
	}

	// Get role credentials from founded user.
	credentials, err := utils.GetCredentialsByRole(foundedUser.UserRole)
	if err != nil {
		// Return status 400 and error message.
        return fiber.StatusBadRequest, "", err, nil
	}

	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewTokens(foundedUser.ID.String(), credentials)
	if err != nil {
		// Return status 500 and token generation error.
        return fiber.StatusInternalServerError, "", err, nil
	}

	// Define user ID.
	userID := foundedUser.ID.String()

	// Create a new Redis connection.
	connRedis, err := cache.RedisConnection()
	if err != nil {
		// Return status 500 and Redis connection error.
        return fiber.StatusInternalServerError, "", err, nil
	}

	// Save refresh token to Redis.
	errSaveToRedis := connRedis.Set(context.Background(), userID, tokens.Refresh, 0).Err()
	if errSaveToRedis != nil {
		// Return status 500 and Redis connection error.
        return fiber.StatusInternalServerError, "", errSaveToRedis, nil
	}

    return fiber.StatusOK, "", nil, tokens
}

func UserSignOut(userID string) (statusCode int, err error) {

	// Create a new Redis connection.
	connRedis, err := cache.RedisConnection()
	if err != nil {
		// Return status 500 and Redis connection error.
        return fiber.StatusInternalServerError, err
	}

	// Save refresh token to Redis.
	errDelFromRedis := connRedis.Del(context.Background(), userID).Err()
	if errDelFromRedis != nil {
		// Return status 500 and Redis deletion error.
        return fiber.StatusInternalServerError, errDelFromRedis
	}
    return fiber.StatusOK, nil
}

