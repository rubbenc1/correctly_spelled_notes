package auth

import "errors"

// User structure
type User struct {
	ID    int
	Token string
}


var users = map[string]User{
	"token1": {ID: 1, Token: "token1"},
	"token2": {ID: 2, Token: "token2"},
}

// AuthenticateUser verifies the token and returns the corresponding user ID
func AuthenticateUser(token string) (int, error) {
	user, exists := users[token]
	if !exists {
		return 0, errors.New("invalid token")
	}
	return user.ID, nil
}
