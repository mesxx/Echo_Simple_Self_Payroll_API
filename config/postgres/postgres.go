package postgres

import (
	"os"

	"github.com/mesxx/Echo_Simple_Self_Payroll_API/model"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	connection := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Error().Msgf("cant connect to database %s", err)
	}
	db.AutoMigrate(&model.Position{}, &model.User{}, &model.Company{}, &model.Transaction{})

	return db

}
