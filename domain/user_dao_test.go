package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	// Shouldnt have a user with id 0, should return nil
	assert.Nil(t, user, "Id 0 not expected")
	assert.NotNil(t, err, "Expected error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User 0 not found", err.Message)

	// tests no assert lib, redundant
	if user != nil {
		t.Error("Id 0 not expected")
	}

	if err == nil {
		t.Error("Expected error when user id is 0")
	}

	// The status code on the error returned should be 404
	if err.StatusCode != http.StatusNotFound {
		t.Error("Expecting 404 when user not found")
	}
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "John", user.FirstName)
	assert.EqualValues(t, "Smith", user.LastName)
	assert.EqualValues(t, "hellO@gmail.com", user.Email)
}
