package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/linuxandchill/mvc/services"
)

// GetUser handles reqs to /user
// ONly responsibiltiy is to get variables needed to handle request
// if there are errors, send them to client
// if all data is there, call services that returns same repsonse from domain
func GetUser(resp http.ResponseWriter, req *http.Request) {

	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("user_id must be a number"))
		// return the bad request to the client
		return
	}

	user, err := services.GetUser(userID)

	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		// Handle the error and return
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
