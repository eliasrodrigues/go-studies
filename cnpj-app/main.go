package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", homeHandler)

	// Serve Static Files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static", http.StripPrefix("/static/", fs))

	log.Println("Server is running on  http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}
