package middleware

import (
	fiber "github.com/gofiber/fiber/v3"
	auth "github.com/CSVaishakh/QuickHand/src/packages/auth/src"
)

func RequireAuth(
	authService *auth.AuthService,
) fiber.Handler {
	return func(c fiber.Ctx) error {
		token := c.Get("Authorization")

		session, claims, err := authService.VerifySession(
			auth.VerifySessionReq{
				Token: token,
			},
		)

		if err != nil {
			return fiber.ErrUnauthorized
		}

		c.Locals("session", session)
		c.Locals("claims", claims)

		return c.Next()
	}
}	