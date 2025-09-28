package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/prankevich/MyProject/internal/config"
	"github.com/prankevich/MyProject/internal/controller"
	"github.com/prankevich/MyProject/internal/repository"
	"github.com/prankevich/MyProject/internal/service"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"log"
	"os"
)

// @title MyProject API
// @contact.name MyProject API Service
// @contact.url http://test.com
// @contact.email test@test.com
func main() {
	logger := Logger()
	if err := config.ReadSettings(); err != nil {
		logger.Error().Err(err).Msg("Error during reading settings")

		return
	}
	logger.Info().Msg("Read settings successfully")
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.AppSettings.PostgresParams.Host,
		config.AppSettings.PostgresParams.Port,
		config.AppSettings.PostgresParams.User,
		os.Getenv("PASWSWORD_DB"),
		config.AppSettings.PostgresParams.Database,
	)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {

		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.AppSettings.RedisParams.Host, config.AppSettings.RedisParams.Port),
		DB:       config.AppSettings.RedisParams.DB,
		Password: "",
	})
	cache := repository.NewCache(rdb, logger)
	repo := repository.NewRepository(db, logger)
	ser := service.NewService(repo, cache, logger)
	ctrl := controller.New(ser, logger)
	ctrl.RunServer(fmt.Sprintf("%s%s", ":", config.AppSettings.AppParams.PortRun))
	if err = db.Close(); err != nil {
		logger.Error().Err(err).Msg("Error during clousing database conectoon" + " Server")

	}
}
func Logger() zerolog.Logger {
	return zerolog.New(os.Stdout).With().Timestamp().Logger()

}
