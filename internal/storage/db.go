package storage

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"shortly/internal/models"
)

var (
	DB    *gorm.DB
	UseDB = true // 👈 asegúrate de que esté exportado (mayúscula)
)

func InitDB() {
	dsn := os.Getenv("DB_URL") // Ej: postgres://user:pass@localhost:5432/shortly
	if dsn == "" {
		log.Println("⚠️ DB_URL no definido, se usará almacenamiento en memoria")
		UseDB = false
		return
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("⚠️ Error al conectar a PostgreSQL: %v\n", err)
		log.Println("➡️  Cambiando a modo memoria")
		UseDB = false
		return
	}

	if err = DB.AutoMigrate(&models.User{}, &models.URL{}); err != nil {
		log.Printf("⚠️ Error al migrar modelos: %v\n", err)
		log.Println("➡️  Cambiando a modo memoria")
		UseDB = false
		return
	}

	log.Println("✅ Base de datos conectada correctamente")
}
