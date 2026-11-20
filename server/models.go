package main

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email        string `gorm"size:100;not null"`
	PasswordHash string `gorm"not null"`
}

type Students struct {
	gorm.Model
	Name       string    `gorm"size:100;not null"`
	Surname    string    `gorm"size:100;not null"`
	Patronymic string    `gorm"size:100;not null"`
	Group      string    `gorm"size:100;not null"`
	BirthDate  time.Time `gorm"time:date;not null"`
}

type Lessons struct {
	gorm.Model
	Subject string    `gorm"size:100;not null"`
	Date    time.Time `gorm"time:date;not null"`
}
