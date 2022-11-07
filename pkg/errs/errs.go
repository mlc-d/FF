package errs

import "errors"

var (
	ErrWrongPassword = errors.New("invalid credentials") // for security reasons, it's better to not disclose wether or not the nick exists in the database
	ErrNickLength    = errors.New("nick length out of bounds")
	ErrInvalidNick   = errors.New("invalid character(s)")

	ErrInvalidFileFormat = errors.New("invalid file format")

	ErrInvalidPayload = errors.New("invalid payload (values cannot be null)")
)
