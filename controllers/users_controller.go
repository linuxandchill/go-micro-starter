package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/linuxandchill/mvc/services"
	"github.com/linuxandchill/mvc/utils"
)

// GetUser handles reqs to /user
// ONly responsibiltiy is to get variables needed to handle request
// if there are errors, send them to client
// if all data is there, call services that returns same repsonse from domain
func GetUser(resp http.ResponseWriter, req *http.Request) {

	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 0, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "userID must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)

		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		// return the bad request to the client
		return
	}

	user, apiErr := services.GetUser(userID)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		// Handle the error and return
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
