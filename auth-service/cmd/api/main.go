package main

import (
	"account-service/config"
	"account-service/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	app := fiber.New()
	config.LoadConfig()

	log.Info().Str("port", config.Config.ApiInfo.Port).Str("host", config.Config.ApiInfo.Host).Msg("Server started")

	db := database.NewPostgres()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	OtpRoutesRegister(app, db)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
		return
	}
}
