package addressService

import "errors"

var(
	ErrAddressNotFoundForUser = errors.New("This address does not belong to this user.")
)