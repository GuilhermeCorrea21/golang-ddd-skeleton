package repository

import (
	"architecture/internal/domain/model"
	user_error "architecture/internal/domain/user/error"

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

func (r *Repository) FindUserByID(id string) (*model.User, error) {
	var user model.User

	result := r.db.Where("id = ?", id).First(&user)
	if result.RowsAffected == 0 {
		return nil, user_error.UserNotFound()
	}

	return &user, nil
}

func (r *Repository) GetUsers() *[]model.User {
	var users []model.User

	r.db.Find(&users)

	return &users
}

func (r *Repository) Create(newUser model.User) (*model.User, error) {
	if err := r.db.Create(&newUser).Error; err != nil {
		return nil, user_error.CreatingUser()
	}

	return &newUser, nil
}
