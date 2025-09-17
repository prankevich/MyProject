package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/prankevich/MyProject/controller"
	"github.com/prankevich/MyProject/repository"
	"github.com/prankevich/MyProject/service"
	"log"
)

func main() {
	dsn := "host=localhost port=5432 user=postgres password=1234 dbname=home_work_db sslmode=disable"

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(db)
	ser := service.NewService(repo)
	ctrl := controller.New(ser)
	ctrl.RunServer(":7777")
	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
	if err = db.Close(); err != nil {
		log.Fatal("Ошибка при подключении:", err)
	}
}
