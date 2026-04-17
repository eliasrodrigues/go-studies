package routes

import (
	"net/http"

	"company-tracker-api/internal/handlers"
)

func RegisterCompanyRoutes(mux *http.ServeMux, handler *handlers.CompanyHandler) {
	mux.HandleFunc("POST /companies", handler.CreateCompany)
	mux.HandleFunc("GET /companies", handler.ListCompanies)
}
