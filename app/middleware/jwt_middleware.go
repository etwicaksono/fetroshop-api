package middleware

import (
	"strings"

	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type JwtMiddleware struct {
	Config *viper.Viper
}

// NewJwtMiddleware creates a new JwtMiddleware instance.
//
// It takes a pointer to a viper.Viper object as the parameter config and returns a pointer to a JwtMiddleware object.
func JwtMiddlewareProvider(config *viper.Viper) *JwtMiddleware {
	return &JwtMiddleware{
		Config: config,
	}
}

// Authenticate authenticates the request using JWT.
//
// It takes a *fiber.Ctx object as a parameter and returns an error.
func (jwtMiddleware *JwtMiddleware) Authenticate(ctx *fiber.Ctx) error {
	header := ctx.Get("Authorization")

	if header == "" {
		return fiber.ErrUnauthorized
	}

	// Spliting the header
	chunks := strings.Split(header, " ")

	// If header signature is not like `Bearer <token>`, then throw
	// This is also required, otherwise chunks[1] will throw out of bound error
	if len(chunks) < 2 {
		return fiber.ErrUnauthorized
	}

	// Verify the token which is in the chunks
	user, err := jwt.Verify(jwtMiddleware.Config.GetString("security.jwt.tokenKey"), chunks[1])

	if err != nil {
		return fiber.ErrUnauthorized
	}

	ctx.Locals("USER", user.ID)

	return ctx.Next()
}
