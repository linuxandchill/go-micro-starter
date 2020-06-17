package domain

import (
	"fmt"
	"net/http"

	"github.com/linuxandchill/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Tyler", LastName: "Munyon", Email: "hellO@gmail.com"},
	}
)

// GetUser from fake database
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
