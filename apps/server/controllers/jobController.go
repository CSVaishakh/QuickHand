package controllers

import (
	"github.com/CSVaishakh/QuickHand/apps/server/middleware"
	job "github.com/CSVaishakh/QuickHand/apps/server/services/jobService"
	auth "github.com/CSVaishakh/QuickHand/packages/auth"

	"github.com/gofiber/fiber/v3"
)

type JobController struct {
	*Controller
	JobService 		*job.JobService
	AuthService 	*auth.AuthService
}

func NewJobController (
	router 			fiber.Router,
	jobService 		*job.JobService,
	authService 	*auth.AuthService,
) *JobController {
	return &JobController{
		Controller: 	NewController(router),
		JobService: 	jobService,
		AuthService: 	authService,
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
}

func (c *JobController) CreateJob(ctx fiber.Ctx) error {
	var req job.CreateJobReq
	if err := ctx.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}

	res, err := c.JobService.CreateJob(req)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(res)
}