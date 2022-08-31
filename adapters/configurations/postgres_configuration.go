package configurations

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfiguration struct {
	DSN string
	DB  *gorm.DB
}

func NewPostgresConfiguration() *PostgresConfiguration {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DBNAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar no banco")
	}

	return &PostgresConfiguration{DSN: dsn, DB: db}
}

func (c *PostgresConfiguration) Migrate(domain interface{}) error {
	return c.DB.AutoMigrate(domain)
}
