package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	CNPJ         string `json:"cnpj" gorm:"uniqueIndex;size:14;not null"`
	RazaoSocial  string `json:"razao_social" gorm:"not null"`
	NomeFantasia string `json:"nome_fantasia"`
	Email        string `json:"email"`
	Telefone     string `json:"telefone"`
	Municipio    string `json:"municipio"`
	UF           string `json:"uf"`
	Notes        string `json:"notes"`
}
