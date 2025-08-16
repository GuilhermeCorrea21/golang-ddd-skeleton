package service

import (
	"architecture/internal/domain/model"
	user_dto "architecture/internal/domain/user/dto"
)

type (
	Repository interface {
		FindUserByID(id string) (*model.User, error)
		GetUsers() *[]model.User
		Create(user user_dto.CreateUserStruct) (*model.User, error)
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

func (s *Service) FindUserByID(id string) (*model.User, error) {
	return s.repository.FindUserByID(id)
}

func (s *Service) GetUsers() *[]model.User {
	return s.repository.GetUsers()
}

func (s *Service) Create(user user_dto.CreateUserStruct) (*model.User, error) {
	return s.repository.Create(user)
}
