package service

import (
	customer_dto "architecture/internal/domain/customer/dto"
	"architecture/internal/domain/model"
	"fmt"
)

type (
	Repository interface {
		GetCustomers() *[]model.Customer
		CreateCustomer(newCustomer customer_dto.CreateCustomerStruct) (*model.Customer, error)
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
	fmt.Println("newCustomer1: ", newCustomer)
	return s.repository.CreateCustomer(newCustomer)
}
