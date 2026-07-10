package controllers

import (
	"errors"
	"fmt"

	"github.com/CSVaishakh/QuickHand/apps/server/middleware"
	alert "github.com/CSVaishakh/QuickHand/apps/server/services/alertService"
	job "github.com/CSVaishakh/QuickHand/apps/server/services/jobService"
	auth "github.com/CSVaishakh/QuickHand/packages/auth"

	"github.com/gofiber/fiber/v3"
)

type JobController struct {
	*Controller
	JobService 		*job.JobService
	AuthService 	*auth.AuthService
	AlertService 	*alert.AlertService
}

func NewJobController (
	router 			fiber.Router,
	jobService 		*job.JobService,
	authService 	*auth.AuthService,
	alertService 	*alert.AlertService,
) *JobController {
	return &JobController{
		Controller: 	NewController(router),
		JobService: 	jobService,
		AuthService: 	authService,
		AlertService:  alertService,
	}
}

func (c *JobController) RegisterRoutes() {

	jobRouter := c.Router.Group("/job")

	jobRouter.Use(
		middleware.RequireAuthHTTPS(
			c.AuthService,
		),
	)

	jobRouter.Post("/", c.CreateJob)
	jobRouter.Patch("/", c.RequestHandyman)
}

func (c *JobController) CreateJob(ctx fiber.Ctx) error {
	var req job.CreateJobReq
	if err := ctx.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}

	claims, ok := ctx.Locals("claims").(*auth.Claims)
	if !ok {
		return fiber.ErrUnauthorized
	}
	req.ClientID = claims.UserID

	res, err := c.JobService.CreateJob(req)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(res)
}

func (c *JobController) RequestHandyman(ctx fiber.Ctx) error {
	var req job.HandymanServiceReq
	var sendAlertReq alert.SendAlertReq

	if err := ctx.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}

	claims, ok := ctx.Locals("claims").(*auth.Claims)
	if !ok {
		return fiber.ErrUnauthorized
	}
	req.ClientID = claims.UserID

	err := c.JobService.RequestHandymanService(req)

	if err != nil {
		if errors.Is(err, job.ErrInvalidHandyman) || errors.Is(err, job.ErrInvalidJob) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		if errors.Is(err, job.ErrTypeMismatch) {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		if errors.Is(err, job.ErrUnauthorized) {
			return fiber.NewError(fiber.StatusForbidden, err.Error())
		}
		return fiber.ErrInternalServerError
	}

	sendAlertReq.UserID = req.HandymanID
	sendAlertReq.Title = "Service Requested"
	sendAlertReq.Description = "someone has requested your service"

	go func() {
		if err := c.AlertService.SendAlert(sendAlertReq); err != nil {
			fmt.Println(err.Error())
		}
	}()

	return ctx.Status(fiber.StatusOK).JSON("Request Successfull")
}
