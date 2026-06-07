package src


import "errors"

var (
    ErrEmailAlreadyExists = errors.New("email already registered")
    ErrInvalidToken    = errors.New("invalid token")
    ErrSignOutFailed   = errors.New("sign out failed")
    ErrSessionNotFound = errors.New("session not found")
    ErrInvalidSession  = errors.New("invalid session")
    ErrSessionExpired  = errors.New("session expired")
    ErrDBFailed        = errors.New("session expired but database failed")
    ErrInvalidCredentials = errors.New("invalid email or password")
    ErrUserDoesNotExist = errors.New("user with email does not exist")
    ErrInvalidOTP = errors.New("otp is invalid")
)