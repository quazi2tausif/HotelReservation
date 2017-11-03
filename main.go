package HotelReservation

import (
	"html/template"
	"net/http"
)

func init() {
	tpl = template.Must(template.ParseGlob("com/WebPage/*.gohtml"))
	fs := http.FileServer(http.Dir("com"))
	http.Handle("/com/", http.StripPrefix("/com/", fs))
	http.HandleFunc("/", createCookie)
	http.HandleFunc("/welcome", welcome)
}
