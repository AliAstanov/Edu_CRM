package main

import (
	"fmt"
	"log"

	"github.com/AliAstanov/Edu_CRM/api"
	"github.com/AliAstanov/Edu_CRM/config"
	"github.com/AliAstanov/Edu_CRM/pkg/db"
	"github.com/AliAstanov/Edu_CRM/service"
	"github.com/AliAstanov/Edu_CRM/storage"
)

func main() {
	fmt.Println("hello")

	cfg := config.Load()

	db, err := db.ConnToDb(cfg.PgConfig)
	if err != nil {
		log.Println("error on con_to_db:", err)
	}

	storage := storage.NewStorage(db)

	service := service.NewService(storage)

	api.Api(service)

}
