package repository

import (
	customer_dto "architecture/internal/domain/customer/dto"
	"architecture/internal/domain/model"
	"fmt"

	"gorm.io/gorm"
)

type (
	Database interface {
		Where(query interface{}, args ...interface{}) (tx *gorm.DB)
		Model(value interface{}) (tx *gorm.DB)
		Updates(values interface{}) (tx *gorm.DB)
		Create(value interface{}) (tx *gorm.DB)
		First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
		Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	}

	Repository struct {
		db Database
	}
)

func New(db Database) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetCustomers() *[]model.Customer {
	var customer []model.Customer

	r.db.Find(&customer)

	return &customer
}

func (r *Repository) CreateCustomer(newCustomer customer_dto.CreateCustomerStruct) (*model.Customer, error) {
	customer := model.Customer{
		Name:   newCustomer.Name,
		Active: true,
	}
	fmt.Println("customer: ", newCustomer)

	if err := r.db.Create(&customer).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}
