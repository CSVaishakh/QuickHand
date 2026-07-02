package controllers

import "github.com/gofiber/fiber/v3"

type Controller struct {
	Router fiber.Router
}

func NewController (router fiber.Router) *Controller {
	return &Controller{
		Router: router,
	}
}