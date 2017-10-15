package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("SitePage/*.html"))
}

func main() {
	fs := http.FileServer(http.Dir("Public"))
	http.Handle("/Public/", http.StripPrefix("/Public/", fs))
	http.HandleFunc("/", landing)
	http.HandleFunc("/welcome/", welcome)
	http.HandleFunc("/login/", login)
	http.HandleFunc("/registration/", registration)
	http.ListenAndServe(":8070", nil)
}

func landing(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "landing.html", nil)
}

func welcome(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "welcome.html", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func registration(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "registration.html", nil)
}
