package usecase

import (
	"fmt"

	"github.com/NursiNursi/laundry-apps/model"
	"github.com/NursiNursi/laundry-apps/model/dto"
	"github.com/NursiNursi/laundry-apps/repository"
)

type ProductUseCase interface {
	RegisterNewProduct(payload model.Product) error
	FindAllProduct(requesPaging dto.PaginationParam) ([]model.Product, dto.Paging, error)
	FindByIdProduct(id string) (model.Product, error)
	UpdateProduct(payload model.Product) error
	DeleteProduct(id string) error
}

type productUseCase struct {
	repo  repository.ProductRepository
	uomUC UomUseCase
}

// RegisterNewProduct implements ProductUseCase.
func (p *productUseCase) RegisterNewProduct(payload model.Product) error {
	if payload.Name == "" || payload.Price == 0 || payload.Uom.Id == "" {
		return fmt.Errorf("name, price and uomID are required fields")
	}

	// cek uom ada atau tidak
	uom, err := p.uomUC.FindByIdUom(payload.Uom.Id)
	if err != nil {
		return fmt.Errorf("uom with ID %s not found", payload.Uom.Id)
	}

	payload.Uom = uom
	err = p.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to register new product: %v", err)
	}
	return nil
}

// FindAllProduct implements ProductUseCase.
func (p *productUseCase) FindAllProduct(requesPaging dto.PaginationParam) ([]model.Product, dto.Paging, error) {
	return p.repo.Paging(requesPaging)
}

// FindByIdProduct implements ProductUseCase.
func (p *productUseCase) FindByIdProduct(id string) (model.Product, error) {
	return p.repo.Get(id)
}

// UpdateProduct implements ProductUseCase.
func (p *productUseCase) UpdateProduct(payload model.Product) error {
	return p.repo.Update(payload)
}

// DeleteProduct implements ProductUseCase.
func (p *productUseCase) DeleteProduct(id string) error {
	return p.repo.Delete(id)
}

func NewProductUseCase(repo repository.ProductRepository, uomUC UomUseCase) ProductUseCase {
	return &productUseCase{repo: repo, uomUC: uomUC}
}
