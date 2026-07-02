package addressService

import "errors"

var(
	ErrAddressUpdateFailed = errors.New("Address Not Updated")
	ErrAddressNotFoundForUser = errors.New("This address does not belong to this user.")
)