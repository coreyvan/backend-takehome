package app

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Run(log *zap.Logger) error {
	dsn := "host=localhost user=candidate password=password123 dbname=telegraph port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Sugar().Fatalf("opening ORM: %v", err)
	}

	port := os.Args[1]
	srv := NewHTTP(log, port, db)

	return srv.Listen()
}
