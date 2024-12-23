package dataBase

import (
	"log"
	// Для подключения БД
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

// Подключение к базе данных
func InitDB() {
	dsn := "host=localhost user=postgres password=12345 dbname=postgres port=5433 sslmode=disable"
	var err error
	// Обработка ошибки
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных:", err)
	}
	log.Println("Успешно соединено с базой данных!")
}
