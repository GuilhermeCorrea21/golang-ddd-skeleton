package app

import (
	"architecture/internal/domain/model"
)

func (app *APP) Migrate() error {
	err := app.DB.AutoMigrate(
		&model.User{},
		&model.Customer{},
		&model.Invoice{},
	)
	if err != nil {
		return err
	}

	return nil
}
