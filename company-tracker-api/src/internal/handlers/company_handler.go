package handlers

import (
	"encoding/json"
	"net/http"

	"company-tracker-api/internal/models"
	"company-tracker-api/internal/services"
)

type CompanyHandler struct {
	Service *services.CompanyService
}

func NewCompanyHandler(service *services.CompanyService) *CompanyHandler {
	return &CompanyHandler{Service: service}
}

func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company

	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if company.CNPJ == "" || company.RazaoSocial == "" {
		http.Error(w, "cnpj and razao_social are required", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateCompany(&company); err != nil {
		http.Error(w, "failed to create company", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(company)
}

func (h *CompanyHandler) ListCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := h.Service.ListCompanies()
	if err != nil {
		http.Error(w, "failed to list companies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companies)
}
