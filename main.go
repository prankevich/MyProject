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
	"log"
	"os"
)

// @title MyProject API
// @contact.name MyProject API Service
// @contact.url http://test.com
// @contact.email test@test.com
func main() {
	if err := config.ReadSettings(); err != nil {
		log.Fatal(err)
	}
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
	cache := repository.NewCache(rdb)
	repo := repository.NewRepository(db)
	ser := service.NewService(repo, cache)
	ctrl := controller.New(ser)
	ctrl.RunServer(fmt.Sprintf("%s%s", ":", config.AppSettings.AppParams.PortRun))
	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
	if err = db.Close(); err != nil {
		log.Fatal("Ошибка при подключении:", err)
	}
}
