package services

import (
	"github.com/linuxandchill/mvc/domain"
	"github.com/linuxandchill/mvc/utils"
)

// GetUser service
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
