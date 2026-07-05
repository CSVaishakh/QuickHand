package middleware

import (
	"strings" 
	
	auth "github.com/CSVaishakh/QuickHand/src/packages/auth/src"

	fiber "github.com/gofiber/fiber/v3"
	"github.com/gofiber/contrib/v3/websocket"
)

func RequireAuthHTTPS(
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

		c.Locals("session", session)
		c.Locals("claims", claims)

		return c.Next()
	}
}

func RequireAuthWS(
	authService *auth.AuthService,
) fiber.Handler {
	return func(c fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
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

			c.Locals("session", session)
			c.Locals("claims", claims)

			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	}
}