package HotelReservation

import (
	"net/http"
)

func createCookie(w http.ResponseWriter, req *http.Request) {

	if userName == "" {
		http.SetCookie(w, &http.Cookie{
			Name:  "HotelReservation.com",
			Value: "Cookie Is Not Active",
		})
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:  userName,
			Value: "Cookie Is Activated",
		})
	}
	http.Redirect(w, req, "/welcome", http.StatusSeeOther)
}
