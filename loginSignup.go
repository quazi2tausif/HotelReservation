package HotelReservation

import (
	"net/http"
)

func loginSignup(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "login-signup.html", nil)
	if err != nil {
		http.Error(w, "We are not sure what happened!", http.StatusNotFound)
	}
}
