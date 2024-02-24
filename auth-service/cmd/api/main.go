package main

import (
	"account-service/config"
	"account-service/pkg/database"
	"fmt"
	"github.com/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Hook(zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
		if level == zerolog.ErrorLevel {
			e.Str("stack", fmt.Sprintf("%+v", errors.WithStack(errors.New(msg))))
		}
	}))
}

func main() {
	app := fiber.New()
	config.LoadConfig()

	db := database.NewPostgres()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	OtpRoutesRegister(app, db)

	err := app.Listen(fmt.Sprintf(":%s", config.Config.ApiInfo.Port))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
		return
	}
}
