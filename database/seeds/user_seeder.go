package seeds

import (
	"tripstory/internal/core/domain/model"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	bytes, err := bcrypt.GenerateFromPassword([]byte("12345678"), 14)

	if err != nil {
		log.Fatal().Err(err).Msg("Error while hashing password")
	}

	admin := model.User{
		Name:     "Ryuu Admin",
		Email:    "ryuuadmin@gmail.com",
		Password: string(bytes),
	}

	if err := db.FirstOrCreate(&admin, model.User{Email: admin.Email}).Error; err != nil {
		log.Fatal().Err(err).Msg("Error while creating admin user")
	} else {
		log.Info().Msg("Admin user created successfully")
	}
}
