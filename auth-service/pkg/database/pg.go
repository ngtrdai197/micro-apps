package database

import (
	"account-service/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Config.PostgresDSN), &gorm.Config{})
	if err != nil {
		log.Panic().Err(err).Msg("Failed to connect to the database")
	}
	return db
}
