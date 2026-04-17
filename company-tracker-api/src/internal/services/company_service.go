package services

import (
	"company-tracker-api/internal/models"

	"gorm.io/gorm"
)

type CompanyService struct {
	DB *gorm.DB
}

func NewCompanyService(db *gorm.DB) *CompanyService {
	return &CompanyService{DB: db}
}

func (s *CompanyService) CreateCompany(company *models.Company) error {
	return s.DB.Create(company).Error
}

func (s *CompanyService) ListCompanies() ([]models.Company, error) {
	var companies []models.Company
	err := s.DB.Order("id desc").Find(&companies).Error
	return companies, err
}
