package app

import (
	"time"
)

type Equipment struct {
	ID              string    `csv:"id" json:"id"`
	Customer        string    `csv:"customer" json:"customer"`
	Fleet           string    `csv:"fleet" json:"fleet"`
	EquipmentID     string    `csv:"equipment_id" json:"equipment_id"`
	EquipmentStatus string    `csv:"equipment_status" json:"equipment_status"`
	DateAdded       time.Time `csv:"date_added" json:"date_added"`
	DateRemoved     time.Time `csv:"date_removed" json:"date_removed,omitempty"`
}

type Location struct {
	ID        string  `csv:"id" json:"id"`
	City      string  `csv:"city" json:"city"`
	CityLong  string  `csv:"city_long" json:"city_long"`
	Station   string  `csv:"station" json:"station"`
	FSAC      string  `csv:"fsac" json:"fsac"`
	SCAC      string  `csv:"scac" json:"scac"`
	SPLC      string  `csv:"splc" json:"splc"`
	State     string  `csv:"state" json:"state"`
	Timezone  string  `csv:"time_zone" json:"time_zone"`
	Longitude float64 `csv:"longitude" json:"longitude"`
	Latitude  float64 `csv:"latitude" json:"latitude"`
	Country   string  `csv:"country" json:"country"`
}

type Waybill struct {
	ID                   string    `csv:"id" json:"id"`
	EquipmentID          string    `csv:"equipment_id" json:"equipment_id"`
	WaybillDate          time.Time `csv:"waybill_date" json:"waybill_date"`
	WaybillNumber        string    `csv:"waybill_number" json:"waybill_number"`
	CreatedDate          time.Time `csv:"created_date" json:"created_date"`
	BillingRoadMarkName  string    `csv:"billing_road_mark_name" json:"billing_road_mark_name"`
	WaybillSourceCode    string    `csv:"waybill_source_code" json:"waybill_source_code"`
	LoadEmptyStatus      string    `csv:"load_empty_status" json:"load_empty_status"`
	OriginMarkName       string    `csv:"origin_mark_name" json:"origin_mark_name"`
	DestinationMarkName  string    `csv:"destination_mark_name" json:"destination_mark_name"`
	SendingRoadMark      string    `csv:"sending_road_mark" json:"sending_road_mark"`
	BillOfLadingNumber   string    `csv:"bill_of_lading_number" json:"bill_of_lading_number"`
	BillOfLadingDate     time.Time `csv:"bill_of_lading_date" json:"bill_of_lading_date"`
	EquipmentWeight      int64     `csv:"equipment_weight" json:"equipment_weight"`
	TareWeight           int64     `csv:"tare_weight" json:"tare_weight"`
	AllowableWeight      int64     `csv:"allowable_weight" json:"allowable_weight"`
	DunnageWeight        int64     `csv:"dunnage_weight" json:"dunnage_weight"`
	EquipmentWeightCode  string    `csv:"equipment_weight_code" json:"equipment_weight_code"`
	CommodityCode        string    `csv:"commodity_code" json:"commodity_code"`
	CommodityDescription string    `csv:"commodity_description" json:"commodity_description"`
	OriginID             string    `csv:"origin_id" json:"origin_id"`
	DestinationID        string    `csv:"destination_id" json:"destination_id"`
	Routes               string    `csv:"routes" json:"routes"`
	Parties              string    `csv:"parties" json:"parties"`
}

type Event struct {
	ID                    string    `csv:"id" json:"id"`
	EquipmentID           string    `csv:"equipment_id" json:"equipment_id"`
	SightingDate          time.Time `csv:"sighting_date" json:"sighting_date"`
	SightingEventCode     string    `csv:"sighting_event_code" json:"sighting_event_code"`
	ReportingRailroadSCAC string    `csv:"reporting_railroad_scac" json:"reporting_railroad_scac"`
	PostingDate           time.Time `csv:"posting_date" json:"posting_date"`
	FromMarkID            string    `csv:"from_mark_id" json:"from_mark_id"`
	LoadEmptyStatus       string    `csv:"load_empty_status" json:"load_empty_status"`
	SightingClaimCode     string    `csv:"sighting_claim_code" json:"sighting_claim_code"`
	SightingEventCodeText string    `csv:"sighting_event_code_text" json:"sighting_event_code_text"`
	TrainID               string    `csv:"train_id" json:"train_id"`
	TrainAlphaCode        string    `csv:"train_alpha_code" json:"train_alpha_code"`
	LocationID            string    `csv:"location_id" json:"location_id"`
	WaybillID             string    `csv:"waybill_id" json:"waybill_id"`
}

type RoutePart struct {
	Scac     string `json:"scac"`
	Junction string `json:"junction,omitempty"`
}
