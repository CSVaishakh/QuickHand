package jobService

import "errors"

var (
	ErrInvalidHandyman = errors.New("Handyman does not exist")
	ErrInvalidJob      = errors.New("Job does not exist, Invalid Job ID")
	ErrTypeMismatch    = errors.New("Job type and Handyman type do not match")
	ErrAlertFailed     = errors.New("Failed to send alert, request sent")
	ErrUnauthorized    = errors.New("This job is not created by you")
)
