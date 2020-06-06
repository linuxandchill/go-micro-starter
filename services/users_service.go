package services

import "github.com/linuxandchill/mvc/domain"

// GetUser service
func GetUser(userID int64) (*domain.User, error) {
	return domain.GetUser(userID)
}
