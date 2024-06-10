package middleware

import (
	"strings"

	"slices"

	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/gofiber/fiber/v2"
)

func ContentTypeMiddleware(c *fiber.Ctx) error {
	contentType := c.Get("Content-Type")
	allowedMethods := []string{"DELETE"}

	if contentType == "" && !slices.Contains(allowedMethods, c.Method()) {
		return errorhelper.Error400("Content-Type header is missing")
	}

	// Check if the Content-Type is application/json
	if !slices.Contains(allowedMethods, c.Method()) &&
		!strings.Contains(contentType, "application/json") &&
		!strings.Contains(contentType, "application/x-www-form-urlencoded") &&
		!strings.Contains(contentType, "multipart/form-data") {
		return errorhelper.ErrorCustom(fiber.StatusUnsupportedMediaType, "Unsupported Content-Type. Only application/json, application/x-www-form-urlencoded and multipart/form-data are allowed")
	}

	// If everything is fine, move on to the next middleware or route handler
	return c.Next()
}
