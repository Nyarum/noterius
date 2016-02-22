package support

import "errors"

type CustomError struct {
	Code uint16
	Err  error
}

func (c *CustomError) Error() string {
	return c.Err.Error()
}

var (
	PlayerIsNotFound CustomError = CustomError{
		1001,
		errors.New("Sorry, player is not found in database"),
	}
	PlayerInGame CustomError = CustomError{
		1104,
		errors.New("Sorry, please wait before auth in this account"),
	}
	PasswordIncorrect CustomError = CustomError{
		1002,
		errors.New("Sorry, your password is incorrect"),
	}
	ClientVersionMismatch CustomError = CustomError{
		7,
		errors.New("Sorry, your client version is mismatched"),
	}
	UnknownError CustomError = CustomError{
		1000,
		errors.New("Sorry, I don't know your problem"),
	}
	SecretPasswordIncorrect CustomError = CustomError{
		534,
		errors.New("Sorry, your secret password is incorrect"),
	}
)
