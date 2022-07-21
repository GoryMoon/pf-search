package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"pfspell-api/internal/server"
	"strconv"
	"syscall"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	level, err := zerolog.ParseLevel(GetEnv("LOG_LEVEL", "info"))
	if err != nil {
		log.Fatal().Str("ctx", "main").Err(err).Msg("Invalid log level")
	}
	log.Logger = log.Logger.Level(level)

	port, err := strconv.Atoi(GetEnv("PORT", "50051"))
	if err != nil {
		log.Fatal().Str("ctx", "main").Err(err).Msg("PORT must be a number")
	}

	meiliHost := GetEnv("MEILI_HOST", "")
	meiliKey := GetEnv("MEILI_KEY", "")
	if len(meiliHost) <= 0 {
		log.Fatal().Str("ctx", "main").Msg("MEILI_HOST must be set")
	}
	if len(meiliKey) <= 0 {
		log.Fatal().Str("ctx", "main").Msg("MEILI_KEY must be set")
	}

	srv := server.CreateNewServer(GetEnv("HOST", ""), port, meiliHost, meiliKey)
	defer srv.Shutdown()

	// Setups a signal to listen for termination
	go func() {
		signChan := make(chan os.Signal, 1)

		signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
		sig := <-signChan
		log.Info().Str("ctx", "main").Str("sig", sig.String()).Msg("Shutting down")

		srv.Shutdown()
	}()

	if err := srv.Run(); err != nil {
		log.Fatal().Str("ctx", "main").Err(err).Send()
	}
	log.Info().Str("ctx", "main").Msg("Goodbye!")
}

func GetEnv(key string, fallback string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}
	return fallback
}
