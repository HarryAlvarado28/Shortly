package storage

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"shortly/internal/models"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DB_URL") // Ejemplo: postgres://user:pass@localhost:5432/shortly
	if dsn == "" {
		log.Fatal("DB_URL no definido en variables de entorno")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a PostgreSQL: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.URL{})
	if err != nil {
		log.Fatal("Error al migrar modelo:", err)
	}
}
