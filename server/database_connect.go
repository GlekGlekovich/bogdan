package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func database_connect() {
	pg_host := os.Getenv("PG_HOST")
	pg_port := os.Getenv("PG_PORT")
	pg_database := os.Getenv("PG_DATABASE")
	pg_user := os.Getenv("PG_USER")
	pg_password := os.Getenv("PG_PASSWORD")

	databaseSourseName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", pg_host, pg_port, pg_database, pg_user, pg_password)

	db, err := gorm.Open(postgres.Open(databaseSourseName, &gorm.Config{}))
	if err != nil {
		log.Fatal("Ошибка подключения к бд", err)
	}

	standartModuleDbStruct, err := db.DB()
	if err != nil {
		return err
	}

}
