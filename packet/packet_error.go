package packet

import "errors"

var (
	PlayerIsNotFound  error = errors.New("Sorry, player is not found in database")
	PlayerInGame      error = errors.New("Sorry, please wait before auth in this account")
	PasswordIncorrect error = errors.New("Sorry, your password is incorrect")
)
