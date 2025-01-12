package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
)

type serviceSale struct {
	Repository repository.SaleRepository
}

func NewServiceSale(sale repository.SaleRepository) *serviceSale {
	return &serviceSale{Repository: sale}
}

func (s *serviceSale) Create(input *domain.SaleInput) (*models.ModelSale, error) {
	var sale domain.SaleInput

	sale.Name = input.Name
	sale.Telepon = input.Telepon
	sale.Email = input.Email
	sale.Alamat = input.Alamat

	res, err := s.Repository.Create(&sale)
	return res, err

}

func (s *serviceSale) Results() (*[]models.ModelSale, error) {
	res, err := s.Repository.Results()
	return res, err
}

func (s *serviceSale) Result(input *domain.SaleInput) (*models.ModelSale, error) {
	var sale domain.SaleInput

	sale.ID = input.ID

	res, err := s.Repository.Result(&sale)

	return res, err
}

func (s *serviceSale) Delete(input *domain.SaleInput) (*models.ModelSale, error) {
	var sale domain.SaleInput

	sale.ID = input.ID

	res, err := s.Repository.Delete(&sale)

	return res, err
}

func (s *serviceSale) Update(input *domain.SaleInput) (*models.ModelSale, error) {
	var sale domain.SaleInput
	sale.Name = input.Name
	sale.Telepon = input.Telepon
	sale.Email = input.Email
	sale.Alamat = input.Alamat

	res, err := s.Repository.Update(&sale)

	return res, err
}
