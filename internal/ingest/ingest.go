package ingest

import (
	"encoding/csv"
	"fmt"
	"github.com/coreyvan/backend-takehome/internal/app"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type Ingester struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewIngester(db *gorm.DB, log *zap.Logger) *Ingester {
	return &Ingester{db: db, log: log}
}

func (i *Ingester) ProcessEvents(filename string) (int, error) {
	if err := i.db.AutoMigrate(&app.Event{}); err != nil {
		return 0, fmt.Errorf("migrating events: %w", err)
	}
	i.db.Exec("DELETE FROM events")

	//	parse from csv
	f, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	lines, err := parseCSVLines(f)
	if err != nil {
		return 0, fmt.Errorf("parsing event lines: %w", err)
	}

	var toSave []app.Event
	for k, l := range lines {
		e, err := parseEvent(l)
		if err != nil {
			i.log.Sugar().Errorf("skipping line %d due to error: %v", k, err)
			continue
		}
		toSave = append(toSave, *e)
	}

	i.db.Create(&toSave)

	return len(toSave), nil
}
func (i *Ingester) ProcessLocations(filename string) (int, error) {
	panic("implement me")
}
func (i *Ingester) ProcessEquipment(filename string) (int, error) {
	panic("implement me")
}
func (i *Ingester) ProcessWaybills(filename string) (int, error) {
	panic("implement me")
}

func parseCSVLines(f *os.File) ([][]string, error) {
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parsing csv: %w", err)
	}

	return lines[1:], nil
}

func parseEvent(line []string) (*app.Event, error) {
	sightingDate, err := time.Parse(timeFormat, line[2])
	if err != nil {
		return nil, fmt.Errorf("parsing sighting date: %w", err)
	}

	postingDate, err := time.Parse(timeFormat, line[5])
	if err != nil {
		return nil, fmt.Errorf("parsing sighting date: %w", err)
	}

	return &app.Event{
		ID:                    line[0],
		EquipmentID:           line[1],
		SightingDate:          sightingDate,
		SightingEventCode:     line[3],
		ReportingRailroadSCAC: line[4],
		PostingDate:           postingDate,
		FromMarkID:            line[6],
		LoadEmptyStatus:       line[7],
		SightingClaimCode:     line[8],
		SightingEventCodeText: line[9],
		TrainID:               line[10],
		TrainAlphaCode:        line[11],
		LocationID:            line[12],
		WaybillID:             line[13],
	}, nil
}
