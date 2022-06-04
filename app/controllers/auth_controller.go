package controllers

import (
	"github.com/0xsp4c3/core/app/models"
	"github.com/0xsp4c3/core/app/services"
	"github.com/0xsp4c3/core/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// UserSignUp method to create a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Param user_role body string true "User role"
// @Success 200 {object} models.User
// @Router /v1/user/sign/up [post]
func UserSignUp(c *fiber.Ctx) error {
	// Create a new user auth struct.
	signUp := &models.SignUp{}

	// Checking received data from JSON body.
	if err := c.BodyParser(signUp); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

    statusCode, message, err, user := services.UserSignUp(signUp)
    if err != nil {
        var msg string
        if message == "" {
            msg = err.Error()
        } else {
            msg = message
        }
        return c.Status(statusCode).JSON(fiber.Map{
            "error": true,
            "msg":   msg,
        })
    }

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}

// UserSignIn method to auth user and return access and refresh tokens.
// @Description Auth user and return access and refresh token.
// @Summary auth user and return access and refresh token
// @Tags User
// @Accept json
// @Produce json
 
// @Param password body string true "User Password"
// @Success 200 {string} status "ok"
// @Router /v1/user/sign/in [post]
func UserSignIn(c *fiber.Ctx) error {
	// Create a new user auth struct.
	signIn := &models.SignIn{}

	// Checking received data from JSON body.
	if err := c.BodyParser(signIn); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

    statusCode, message, err, tokens := services.UserSignIn(signIn)
    if err != nil {
        var msg string
        if message == "" {
            msg = err.Error()
        } else {
            msg = message
        }
        return c.Status(statusCode).JSON(fiber.Map{
            "error": true,
            "msg":   msg,
        })
    }

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"tokens": fiber.Map{
			"access":  tokens.Access,
			"refresh": tokens.Refresh,
		},
	})
}

// UserSignOut method to de-authorize user and delete refresh token from Redis.
// @Description De-authorize user and delete refresh token from Redis.
// @Summary de-authorize user and delete refresh token from Redis
// @Tags User
// @Accept json
// @Produce json
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/user/sign/out [post]
func UserSignOut(c *fiber.Ctx) error {
	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Define user ID.
	userID := claims.UserID.String()

    statusCode, err := services.UserSignOut(userID)
    if err != nil {
        return c.Status(statusCode).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
