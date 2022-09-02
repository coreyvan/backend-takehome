package main

import (
	"fmt"
	"github.com/coreyvan/backend-takehome/internal/ingest"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	if err := run(log); err != nil {
		log.Sugar().Fatal(err)
	}
}

func run(log *zap.Logger) error {
	command := os.Args[1]

	switch command {
	case "ingest":
		dsn := "host=localhost user=candidate password=password123 dbname=telegraph port=5432 sslmode=disable"
		db, err := gorm.Open(postgres.Open(dsn))
		if err != nil {
			return fmt.Errorf("opening db: %w", err)
		}
		i := ingest.NewIngester(db, log)
		kind := os.Args[2]
		switch kind {
		case "locations":
			log.Sugar().Infof("ingesting locations...")
			n, err := i.ProcessLocations("data/locations.csv")
			if err != nil {
				return fmt.Errorf("processing locations: %w", err)
			}
			log.Sugar().Infof("success... ingested %d rows", n)
		case "equipment":
			log.Sugar().Infof("ingesting equipment...")
			n, err := i.ProcessEquipment("data/equipment.csv")
			if err != nil {
				return fmt.Errorf("processing equipment: %w", err)
			}
			log.Sugar().Infof("success... ingested %d rows", n)
		case "events":
			log.Sugar().Infof("ingesting events...")
			n, err := i.ProcessEvents("data/events.csv")
			if err != nil {
				return fmt.Errorf("processing events: %w", err)
			}
			log.Sugar().Infof("success... ingested %d rows", n)
		case "waybills":
			log.Sugar().Infof("ingesting waybills...")
			n, err := i.ProcessWaybills("data/waybills.csv")
			if err != nil {
				return fmt.Errorf("processing waybills: %w", err)
			}
			log.Sugar().Infof("success... ingested %d rows", n)
		default:
			return fmt.Errorf("invalid kind %s", kind)
		}
	default:
		return fmt.Errorf("invalid command %s", command)
	}

	return nil
}
