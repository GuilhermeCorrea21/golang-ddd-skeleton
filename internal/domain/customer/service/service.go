package service

import (
	customer_dto "architecture/internal/domain/customer/dto"
	"architecture/internal/domain/model"
)

type (
	Repository interface {
		GetCustomers() *[]model.Customer
		CreateCustomer(newCustomer model.Customer) (*model.Customer, error)
	}

	Service struct {
		repository Repository
	}
)

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetCustomers() *[]model.Customer {
	return s.repository.GetCustomers()
}

func (s *Service) CreateCustomer(newCustomer customer_dto.CreateCustomerStruct) (*model.Customer, error) {
	customer := model.Customer{
		Name:   newCustomer.Name,
		Active: true,
	}

	resp, err := s.repository.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
