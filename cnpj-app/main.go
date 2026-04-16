package main

import (
	"html/template"
	"log"
	"net/http"

	"cnpj-app/handlers"
)

func main() {
	templates := template.Must(template.ParseGlob("templates/*.html"))

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.HomeHandler(templates))
	http.HandleFunc("/buscar", handlers.BuscarHandler(templates))

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
