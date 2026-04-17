package main

import (
	"log"
	"net/http"

	"company-tracker-api/internal/database"
	"company-tracker-api/internal/handlers"
	"company-tracker-api/internal/routes"
	"company-tracker-api/internal/services"
)

func main() {
	db := database.Connect()

	companyService := services.NewCompanyService(db)
	companyHandler := handlers.NewCompanyHandler(companyService)

	mux := http.NewServeMux()
	routes.RegisterCompanyRoutes(mux, companyHandler)

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
