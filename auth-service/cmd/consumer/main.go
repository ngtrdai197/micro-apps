package main

import (
	"account-service/config"
	"account-service/pkg/kafka"
	"fmt"
	"github.com/pkg/errors"
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
	config.LoadConfig()

	c := kafka.NewConsumer()
	c.StartConsuming()
}
