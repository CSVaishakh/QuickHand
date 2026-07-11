package alertService
import (
	"github.com/google/uuid"
)

type SendAlertReq struct {
	UserID 			uuid.UUID
	Title				string
	Description 	string
}

type AlertPayload struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}