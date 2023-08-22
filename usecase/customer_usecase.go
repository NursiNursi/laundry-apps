package usecase

import (
	"fmt"

	"github.com/NursiNursi/laundry-apps/model"
	"github.com/NursiNursi/laundry-apps/model/dto"
	"github.com/NursiNursi/laundry-apps/repository"
)

type CustomerUseCase interface {
	RegisterNewCustomer(payload model.Customer) error
	FindAllCustomer(requesPaging dto.PaginationParam) ([]model.Customer, dto.Paging, error)
	FindByIdCustomer(id string) (model.Customer, error)
	UpdateCustomer(payload model.Customer) error
	DeleteCustomer(id string) error
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

// DeleteCustomer implements CustomerUseCase.
func (c *customerUseCase) DeleteCustomer(id string) error {
	customer, err := c.FindByIdCustomer(id)
	if err != nil {
		return fmt.Errorf("customer with ID %s not found", id)
	}

	err = c.repo.Delete(customer.Id)
	if err != nil {
		return fmt.Errorf("failed to delete customer: %v", err.Error())
	}
	return nil
}

// FindAllProduct implements CustomerUseCase.
func (c *customerUseCase) FindAllCustomer(requesPaging dto.PaginationParam) ([]model.Customer, dto.Paging, error) {
	return c.repo.Paging(requesPaging)
}

// FindByIdCustomer implements CustomerUseCase.
func (c *customerUseCase) FindByIdCustomer(id string) (model.Customer, error) {
	return c.repo.Get(id)
}

// RegisterNewCustomer implements CustomerUseCase.
func (c *customerUseCase) RegisterNewCustomer(payload model.Customer) error {
	if payload.Name == "" || payload.PhoneNumber == "" {
		return fmt.Errorf("name, phone number are required fields")
	}
	customer, _ := c.repo.GetPhoneNumber(payload.PhoneNumber)
	if customer.PhoneNumber == payload.PhoneNumber {
		return fmt.Errorf("customer with phone number %s already exists", payload.PhoneNumber)
	}
	err := c.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create customer: %v", err.Error())
	}
	return nil
}

// UpdateCustomer implements CustomerUseCase.
func (c *customerUseCase) UpdateCustomer(payload model.Customer) error {
	if payload.Name == "" || payload.PhoneNumber == "" {
		return fmt.Errorf("name, phone number are required fields")
	}
	customer, _ := c.repo.GetPhoneNumber(payload.PhoneNumber)
	if customer.PhoneNumber == payload.PhoneNumber && customer.Id != payload.Id {
		return fmt.Errorf("customer with phone number %s already exists", payload.PhoneNumber)
	}
	err := c.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update customer: %v", err.Error())
	}
	return nil
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{repo: repo}
}
