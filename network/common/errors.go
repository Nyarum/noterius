package common

type NetworkError struct {
	ID    uint16
	Value string
}

func (n NetworkError) Error() string {
	return n.Value
}

func (n NetworkError) GetID() uint16 {
	return n.ID
}

func NewNetworkError(id uint16, value string) NetworkError {
	return NetworkError{
		ID:    id,
		Value: value,
	}
}

var (
	PlayerIsNotFound        NetworkError = NewNetworkError(1001, "Player is not found in store")
	PlayerInGame            NetworkError = NewNetworkError(1104, "Player still in game")
	PasswordIncorrect       NetworkError = NewNetworkError(1002, "Password is incorrect")
	ClientVersionMismatch   NetworkError = NewNetworkError(7, "Version is mismatch")
	UnknownError            NetworkError = NewNetworkError(1000, "Unknown error")
	SecretPasswordIncorrect NetworkError = NewNetworkError(534, "Secret password is incorrect")
	ExistCharName           NetworkError = NewNetworkError(526, "Character name is exists")
	InvalidCharName         NetworkError = NewNetworkError(531, "Invalid character name")
	InvalidBirthLocation    NetworkError = NewNetworkError(527, "Map isn't exists")
	InternalError           NetworkError = NewNetworkError(521, "Internal error")
)
