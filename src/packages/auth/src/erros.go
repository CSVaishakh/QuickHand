package src


import "errors"

var (
	ErrEmailAldreadyExists = errors.New("email ardeady registered")
	ErrInvalidToken    = errors.New("invalid token")
	ErrSignOutFailed   = errors.New("sign out failed")
	ErrSessionNotFound = errors.New("session not found")
	ErrInvalidSession  = errors.New("invalid session")
	ErrSessionExpired  = errors.New("session expired")
	ErrDBFailed        = errors.New("session expired but databse failed")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserDoesNotExist = errors.New("user with email dopes not exist")
	ErrInvalidOTP = errors.New("otp is is invalid")
)
