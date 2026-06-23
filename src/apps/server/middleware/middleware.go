package middleware

import (
	"strings" 

	fiber "github.com/gofiber/fiber/v3"
	auth "github.com/CSVaishakh/QuickHand/src/packages/auth/src"
)

func RequireAuth(
	authService *auth.AuthService,
) fiber.Handler {
	return func(c fiber.Ctx) error {
		token := c.Get("Authorization")

		if after, ok := strings.CutPrefix(token, "Bearer "); ok  {
			token = after
		}

		session, claims, err := authService.VerifySession(
			auth.VerifySessionReq{
				Token: token,
			},
		)

		if err != nil {
			return fiber.ErrUnauthorized
		}

		c.Locals("token", token)
		c.Locals("session", session)
		c.Locals("claims", claims)

		return c.Next()
	}
}