package app

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run(log *zap.Logger) error {
	dsn := "host=localhost user=candidate password=password123 dbname=telegraph port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Sugar().Fatalf("opening ORM: %v", err)
	}

	srv := NewHTTP(log, 3000, db)

	return srv.Listen()
}
