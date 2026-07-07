package controllers

import(
	"errors"
	
	"github.com/CSVaishakh/QuickHand/apps/server/services/addressService"
	"github.com/CSVaishakh/QuickHand/apps/server/middleware"
	auth "github.com/CSVaishakh/QuickHand/packages/auth"

	"github.com/gofiber/fiber/v3"
)

type AddressController struct {
	*Controller
	AddressService 	*addressService.AddressService 
	AuthService 		*auth.AuthService
}

func NewAddressController(
	router 				fiber.Router,
	addressService 	*addressService.AddressService,
	authService 		*auth.AuthService,
) *AddressController {
	return &AddressController{
		Controller: 		NewController(router),
		AddressService: 	addressService,
		AuthService: 		authService,	
	}
}

func (c *AddressController) RegisterRoutes() {
	addressRouter := c.Router.Group("/address")

	addressRouter.Use(
		middleware.RequireAuthHTTPS(
			c.AuthService,
		),
	)

	addressRouter.Post("/add-new-address", c.AddNewAddress)
	addressRouter.Put("/update-address", c.UpdateAddress)
	addressRouter.Get("/get-address", c.GetAddress)
}

func (c *AddressController) AddNewAddress(ctx fiber.Ctx) error {
	var req addressService.AddAddressReq

	if err := ctx.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}

	claims, ok := ctx.Locals("claims").(*auth.Claims)
	if !ok {
		return fiber.ErrUnauthorized
	}
	req.UserId = claims.UserID

	res, err := c.AddressService.AddNewAddress(req)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(res)
}

func (c *AddressController) UpdateAddress(ctx fiber.Ctx) error {
	var req addressService.UpdateAddressReq

	if err := ctx.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}

	claims, ok := ctx.Locals("claims").(*auth.Claims)
	if !ok {
		return fiber.ErrUnauthorized
	}
	req.UserId = claims.UserID

	res, err := c.AddressService.UpdateAddress(req)

	if errors.Is(err, addressService.ErrAddressNotFoundForUser) {
		return fiber.NewError(
			fiber.StatusForbidden,
			err.Error(),
		)
	}

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(res)
}

func (c *AddressController) GetAddress(ctx fiber.Ctx) error {
	var req addressService.GetAddressesReq
	
	claims, ok := ctx.Locals("claims").(*auth.Claims)
	if !ok {
		return fiber.ErrUnauthorized
	}
	req.UserId = claims.UserID

	res, err := c.AddressService.GetAddresses(req)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(res)
}