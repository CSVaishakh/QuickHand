package controllers

import (
	"errors"

	"github.com/CSVaishakh/QuickHand/src/apps/server/middleware"
	auth "github.com/CSVaishakh/QuickHand/src/packages/auth/src"

	"github.com/gofiber/fiber/v3"
)

type AuthController struct {
	*Controller
	AuthService *auth.AuthService
}

func NewAuthController(
	router fiber.Router,
	authService *auth.AuthService,
) *AuthController {
	return &AuthController{
		Controller:  NewController(router),
		AuthService: authService,
	}
}

func (c *AuthController) RegisterRoutes() {
	authRouter := c.Router.Group("/auth")

	authRouter.Post("/client/sign-up", c.ClientSignUp)
	authRouter.Post("/client/sign-in", c.ClientSignIn)

	authRouter.Post("/handyman/sign-up", c.HandymanSignUp)
	authRouter.Post("/handyman/sign-in", c.HandymanSignIn)

	authRouter.Use(
		middleware.RequireAuth(
			c.AuthService,
		),
	)

	authRouter.Post("/sign-out", c.SignOut)
}

func (c *AuthController) HandymanSignUp(ctx fiber.Ctx) error {
	var req auth.HandymanSignUpReq

	if err := ctx.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}

	res, err := c.AuthService.HandymanSignUp(req)

	if errors.Is(err, auth.ErrEmailAlreadyExists) {
		return fiber.NewError(
			fiber.StatusConflict,
			err.Error(),
		)
	}

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(res)
}

func (c *AuthController) ClientSignUp(ctx fiber.Ctx) error {
	var req auth.ClientSignUpReq

	if err := ctx.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}

	res, err := c.AuthService.ClientSignUp(req)

	if errors.Is(err, auth.ErrEmailAlreadyExists) {
		return fiber.NewError(
			fiber.StatusConflict,
			err.Error(),
		)
	}

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(res)
}

func (c *AuthController) HandymanSignIn(ctx fiber.Ctx) error {
	var req auth.SignInReq

	if err := ctx.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}

	res, err := c.AuthService.HandymanSignIn(req)

	if errors.Is(err, auth.ErrInvalidCredentials) {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			err.Error(),
		)
	}

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(res)
}

func (c *AuthController) ClientSignIn(ctx fiber.Ctx) error {
	var req auth.SignInReq

	if err := ctx.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}

	res, err := c.AuthService.ClientSignIn(req)

	if errors.Is(err, auth.ErrInvalidCredentials) {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			err.Error(),
		)
	}

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(res)
}

func (c *AuthController) SignOut(ctx fiber.Ctx) error {
	token, ok := ctx.Locals("token").(string)
	if !ok {
		return fiber.ErrUnauthorized
	}

	if err := c.AuthService.SignOut(token); err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}