package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"cnpj-app/models"
)

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return e.Message
}

func BuscarCNPJ(cnpj string) (*models.Empresa, error) {
	url := fmt.Sprintf("https://api.opencnpj.org/%s?datasets=receita", cnpj)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call external API: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "CNPJ not found",
		}
	case http.StatusTooManyRequests:
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "Rate limit exceeded. Try again in a moment",
		}
	default:
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("unexpected API status: %d", resp.StatusCode),
		}
	}

	var empresa models.Empresa

	if err := json.NewDecoder(resp.Body).Decode(&empresa); err != nil {
		return nil, fmt.Errorf("failed to decode API response: %w", err)
	}

	empresa.CNPJ = strings.TrimSpace(empresa.CNPJ)

	return &empresa, nil
}
