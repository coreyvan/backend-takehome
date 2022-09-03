package ingest

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/coreyvan/backend-takehome/internal/app"
	"github.com/gocarina/gocsv"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

var validTimeFormats = []string{time.RFC3339, "2006-01-02 15:04:05", "2006-01-02 15:04:05.000000"}

type Ingester struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewIngester(db *gorm.DB, log *zap.Logger) *Ingester {
	return &Ingester{db: db, log: log}
}

func (i *Ingester) ProcessEvents(filename string) (int, error) {
	i.db.Exec("DROP TABLE IF EXISTS events")
	if err := i.db.AutoMigrate(&app.Event{}); err != nil {
		return 0, fmt.Errorf("migrating events: %w", err)
	}

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
	i.db.Exec("DROP TABLE IF EXISTS locations")
	if err := i.db.AutoMigrate(&app.Location{}); err != nil {
		return 0, fmt.Errorf("migrating locations: %w", err)
	}

	//	parse from csv
	f, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	var toSave []app.Location

	if err := gocsv.UnmarshalFile(f, &toSave); err != nil {
		return 0, fmt.Errorf("unmarshaling file: %w", err)
	}

	i.db.Create(&toSave)

	return len(toSave), nil
}
func (i *Ingester) ProcessEquipment(filename string) (int, error) {
	i.db.Exec("DROP TABLE IF EXISTS equipment")
	if err := i.db.AutoMigrate(&app.Equipment{}); err != nil {
		return 0, fmt.Errorf("migrating equipment: %w", err)
	}

	//	parse from csv
	f, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	lines, err := parseCSVLines(f)
	if err != nil {
		return 0, fmt.Errorf("parsing equipment lines: %w", err)
	}

	var toSave []app.Equipment
	for k, l := range lines {
		e, err := parseEquipment(l)
		if err != nil {
			i.log.Sugar().Errorf("skipping line %d due to error: %v", k, err)
			continue
		}
		toSave = append(toSave, *e)
	}

	i.db.Create(&toSave)

	return len(toSave), nil
}
func (i *Ingester) ProcessWaybills(filename string) (int, error) {
	i.db.Exec("DROP TABLE IF EXISTS waybills")
	if err := i.db.AutoMigrate(&app.Waybill{}); err != nil {
		return 0, fmt.Errorf("migrating waybills: %w", err)
	}

	//	parse from csv
	f, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	lines, err := parseCSVLines(f)
	if err != nil {
		return 0, fmt.Errorf("parsing waybill lines: %w", err)
	}

	var toSave []app.Waybill
	for k, l := range lines {
		w, err := parseWaybill(l)
		if err != nil {
			i.log.Sugar().Errorf("skipping line %d due to error: %v", k, err)
			continue
		}
		toSave = append(toSave, *w)
	}

	i.db.Create(&toSave)

	return len(toSave), nil
}

func parseCSVLines(f *os.File) ([][]string, error) {
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parsing csv: %w", err)
	}

	return lines[1:], nil
}

func parseEvent(line []string) (*app.Event, error) {
	sightingDate, err := parseTime(line[2])
	if err != nil {
		return nil, fmt.Errorf("parsing sighting date: %w", err)
	}

	postingDate, err := parseTime(line[5])
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

func parseWaybill(line []string) (*app.Waybill, error) {
	waybillDate, err := parseTime(line[2])
	if err != nil {
		return nil, fmt.Errorf("parsing sighting date: %w", err)
	}

	createdDate, err := parseTime(line[4])
	if err != nil {
		return nil, fmt.Errorf("parsing created date: %w", err)
	}

	billOfLadingDate, err := parseTime(line[12])
	if err != nil {
		return nil, fmt.Errorf("parsing bill of lading date: %w", err)
	}

	equipmentWeight, err := strconv.ParseInt(line[13], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parsing equipment weight")
	}

	tareWeight, err := strconv.ParseInt(line[14], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parsing equipment weight")
	}

	allowableWeight, err := strconv.ParseInt(line[15], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parsing equipment weight")
	}

	dunnageWeight, err := strconv.ParseInt(line[16], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parsing equipment weight")
	}

	return &app.Waybill{
		ID:                   line[0],
		EquipmentID:          line[1],
		WaybillDate:          waybillDate,
		WaybillNumber:        line[3],
		CreatedDate:          createdDate,
		BillingRoadMarkName:  line[5],
		WaybillSourceCode:    line[6],
		LoadEmptyStatus:      line[7],
		OriginMarkName:       line[8],
		DestinationMarkName:  line[9],
		SendingRoadMark:      line[10],
		BillOfLadingNumber:   line[11],
		BillOfLadingDate:     billOfLadingDate,
		EquipmentWeight:      equipmentWeight,
		TareWeight:           tareWeight,
		AllowableWeight:      allowableWeight,
		DunnageWeight:        dunnageWeight,
		EquipmentWeightCode:  line[17],
		CommodityCode:        line[18],
		CommodityDescription: line[19],
		OriginID:             line[20],
		DestinationID:        line[21],
		Routes:               line[22],
		Parties:              line[23],
	}, nil
}

func parseEquipment(line []string) (*app.Equipment, error) {
	dateAdded, err := parseTime(line[5])
	if err != nil {
		return nil, err
	}

	dateRemoved, err := parseTime(line[6])
	if err != nil {
		return nil, err
	}

	return &app.Equipment{
		ID:              line[0],
		Customer:        line[1],
		Fleet:           line[2],
		EquipmentID:     line[3],
		EquipmentStatus: line[4],
		DateAdded:       dateAdded,
		DateRemoved:     dateRemoved,
	}, nil
}

func parseTime(str string) (time.Time, error) {
	if str == "" {
		return time.Time{}, nil
	}

	for _, f := range validTimeFormats {
		t, err := time.Parse(f, str)
		if err == nil {
			return t.UTC(), nil
		}
	}
	return time.Time{}, errors.New("could not parse string with any valid formats")
}
