package app

import (
	"net/http"

	"github.com/mo-mirzania/test/controller"
)

func StartApplication() {
	http.HandleFunc("/ping", controller.Ping)
	http.HandleFunc("/users", controller.UserID)
	http.ListenAndServe(":8080", nil)
}
