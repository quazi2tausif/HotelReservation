package HotelReservation

import (
	"net/http"
)

func welcome(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie(userName)
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if c.Value == "Cookie Is Not Active" {
		err = tpl.ExecuteTemplate(w, "welcome.html", nil)
		if err != nil {
			http.Error(w, "We are not sure what happened!", http.StatusNotFound)
		}
	} else {
		err = tpl.ExecuteTemplate(w, "welcome.html", c)
		if err != nil {
			http.Error(w, "We are not sure what happened!", http.StatusNotFound)
		}
	}
}
