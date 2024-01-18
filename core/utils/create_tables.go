package utils

import (
	"smartCampus/core/models"
)

func CreateTables() {
	models.DB.Debug().AutoMigrate(
		&models.User{},
		&models.Data{},
	)

}
