package components

import (
	"first-api/src/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=todos port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&models.Todo{}); err != nil {
		log.Println(err)
	}

	return db, err
}
