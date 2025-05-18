package front

import (
	"net/http"
	"text/template"
)

func RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/", Homepage)
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

// Return default Homepage, a simple alpineJS application to users
func Homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	template, err := template.ParseFiles("templates/index.html.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := template.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
