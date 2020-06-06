package domain

import (
	"fmt"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Tyler", LastName: "Munyon", Email: "hellO@gmail.com"},
	}
)

// GetUser from fake database
func GetUser(userID int64) (*User, error) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, fmt.Errorf("User %v not found", userID)

}
