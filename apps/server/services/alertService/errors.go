package alertService

import "errors"

var (
	ErrSavingAlert = errors.New("alert sent! error saving alert to database")
)