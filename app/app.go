package app

import (
	"net/http"

	"github.com/linuxandchill/mvc/controllers"
)

// StartApp defines handlers for the application
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
