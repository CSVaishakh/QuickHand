package websockets

import "errors"

var(
	ErrEmptyUserId = errors.New("Empty UserID")
	ErrEmptyConnectionReference = errors.New("Empty connection refernce")
)