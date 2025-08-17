package repository

import (
	customer_error "architecture/internal/domain/customer/error"
	"architecture/internal/domain/model"

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

func (r *Repository) CreateCustomer(newCustomer model.Customer) (*model.Customer, error) {
	if err := r.db.Create(&newCustomer).Error; err != nil {
		return nil, customer_error.CreatingCustomer()
	}

	return &newCustomer, nil
}
