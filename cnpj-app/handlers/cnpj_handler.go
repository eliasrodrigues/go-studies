package handlers

import (
	"html/template"
	"net/http"
	"regexp"

	"cnpj-app/models"
	"cnpj-app/services"
)

type IndexData struct {
	Empresa *models.Empresa
	Error   string
	CNPJ    string
}

var onlyDigits = regexp.MustCompile(`\D`)

func HomeHandler(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := IndexData{}
		if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
		}
	}
}

func BuscarHandler(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cnpjInput := r.URL.Query().Get("cnpj")
		cnpj := onlyDigits.ReplaceAllString(cnpjInput, "")

		data := IndexData{
			CNPJ: cnpjInput,
		}

		if cnpj == "" {
			data.Error = "Please enter a CNPJ"
			renderTemplate(w, templates, data)
			return
		}

		if len(cnpj) != 14 {
			data.Error = "CNPJ must contain exactly 14 digits"
			renderTemplate(w, templates, data)
			return
		}

		empresa, err := services.BuscarCNPJ(cnpj)
		if err != nil {
			data.Error = err.Error()
			renderTemplate(w, templates, data)
			return
		}

		data.Empresa = empresa
		renderTemplate(w, templates, data)
	}
}

func renderTemplate(w http.ResponseWriter, templates *template.Template, data IndexData) {
	if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "failed to render template", http.StatusInternalServerError)
	}
}
