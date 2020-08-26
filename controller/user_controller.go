package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/mo-mirzania/test/service"
	"github.com/mo-mirzania/test/utils/resterror"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Pong")
}

func UserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 0)
	if err != nil {
		io.WriteString(w, "Invalid user-id")
		return
	}
	user, restErr := service.UserServices.GetUser(int(userID))
	if restErr != nil {
		er := resterror.NotFoundError("Not found")
		erJsoned, _ := json.Marshal(er)
		w.Write(erJsoned)
		return
	}
	jsoned, _ := json.Marshal(user)
	w.Write(jsoned)
}
