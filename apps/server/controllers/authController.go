package controllers

import (
	"errors"
	"strings"

	"github.com/CSVaishakh/QuickHand/apps/server/middleware"
	auth "github.com/CSVaishakh/QuickHand/packages/auth"

	"github.com/gofiber/fiber/v3"	
)

type AuthController struct {
	*Controller
	AuthService *auth.AuthService
}

func NewAuthController(
	router 			fiber.Router,
	authService 	*auth.AuthService,
) *AuthController {
	return &AuthController{
		Controller:  	NewController(router),
		AuthService: 	authService,
	}
}

func (c *AuthController) RegisterRoutes() {
	authRouter := c.Router.Group("/auth")

	authRouter.Post("/client/sign-up", c.ClientSignUp)
	authRouter.Post("/client/sign-in", c.ClientSignIn)

	authRouter.Post("/handyman/sign-up", c.HandymanSignUp)
	authRouter.Post("/handyman/sign-in", c.HandymanSignIn)

	authRouter.Post("/forgot-password", c.ForgotPassword)
	authRouter.Post("/forgot-password/verify-otp", c.VerifyOTP)
	authRouter.Post("/forgot-password/reset", c.ResetPassword)

	authRouter.Use(
		middleware.RequireAuthHTTPS(
			c.AuthService,
		),
	)

	authRouter.Post("/sign-out", c.SignOut)
	authRouter.Get("/session", c.GetSession)
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

func (c *AuthController) ForgotPassword(ctx fiber.Ctx) error {
	var req auth.ForgotPasswordReq

	if err := ctx.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}
	
	token, err := c.AuthService.ForgotPassword(req)

	if errors.Is(err, auth.ErrUserDoesNotExist){
		return fiber.NewError(
			fiber.StatusUnauthorized,
			auth.ErrInvalidCredentials.Error(),
		)
	}

	if err != nil {
		return fiber.ErrInternalServerError
	}

	ctx.Status(fiber.StatusOK)
	ctx.JSON(token)

	return nil
}

func (c *AuthController) VerifyOTP(ctx fiber.Ctx) error {
	var req auth.OtpVerificationReq

	if err := ctx.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}

	token, err := c.AuthService.VerifyOTP(req)
	if errors.Is(err, auth.ErrInvalidOTP){
		return fiber.NewError(
			fiber.ErrUnauthorized.Code,
			auth.ErrInvalidOTP.Error(),
		)
	}

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.Status(fiber.StatusOK).JSON(token)
}

func (c *AuthController) ResetPassword(ctx fiber.Ctx) error {
	var req auth.ResetPasswordReq

	if err := ctx.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}

	err := c.AuthService.ResetPassword(req)

	if err != nil {
		if errors.Is(err, auth.ErrInvalidToken) {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				err.Error(),
			)
		}

		if errors.Is(err, auth.ErrUserDoesNotExist) {
			return fiber.NewError(
				fiber.StatusNotFound,
				err.Error(),
			)
		}

		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c *AuthController) GetSession(ctx fiber.Ctx) error {
	token := ctx.Get("Authorization")

   after, ok := strings.CutPrefix(token, "Bearer ")
   if !ok {
        return fiber.ErrUnauthorized
   }

   req := auth.GetSessionReq{
        Token: after,
   }

	session, err := c.AuthService.GetSession(req)

	switch {
	case errors.Is(err, auth.ErrInvalidToken):
		return fiber.NewError(
			fiber.StatusUnauthorized,
			err.Error(),
		)

	case errors.Is(err, auth.ErrSessionNotFound):
		return fiber.NewError(
			fiber.StatusUnauthorized,
			err.Error(),
		)

	case errors.Is(err, auth.ErrInvalidCredentials):
		return fiber.NewError(
			fiber.StatusUnauthorized,
			err.Error(),
		)

	case err != nil:
		return fiber.ErrInternalServerError
	}

	return ctx.Status(fiber.StatusOK).JSON(session)
}