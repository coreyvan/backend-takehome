package app

type Equipment struct {
	Id              int64  `csv:"id"`
	Customer        string `csv:"customer"`
	Fleet           string `csv:"fleet"`
	EquipmentID     string `csv:"equipment_id"`
	EquipmentStatus bool   `csv:"equipment_status"`
	DateAdded       string `csv:"date_added"`
	DateRemoved     string `csv:"date_removed"`
}

type Location struct {
	Id        bool    `csv:"id"`
	City      string  `csv:"city"`
	CityLong  string  `csv:"city_long"`
	Station   string  `csv:"station"`
	FSAC      int64   `csv:"fsac"`
	SCAC      string  `csv:"scac"`
	SPLC      int64   `csv:"splc"`
	State     string  `csv:"state"`
	Timezone  string  `csv:"time_zone"`
	Longitude float64 `csv:"longitude"`
	Latitude  float64 `csv:"latitude"`
	Country   string  `csv:"country"`
}

type Waybill struct {
	Id                   bool   `csv:"id"`
	EquipmentID          string `csv:"equipment_id"`
	WaybillDate          string `csv:"waybill_date"`
	WaybillNumber        int64  `csv:"waybill_number"`
	CreatedDate          string `csv:"created_date"`
	BillingRoadMarkName  string `csv:"billing_road_mark_name"`
	WaybillSourceCode    int64  `csv:"waybill_source_code"`
	LoadEmptyStatus      string `csv:"load_empty_status"`
	OriginMarkName       string `csv:"origin_mark_name"`
	DestinationMarkName  string `csv:"destination_mark_name"`
	SendingRoadMark      string `csv:"sending_road_mark"`
	BillOfLadingNumber   int64  `csv:"bill_of_lading_number"`
	BillOfLadingDate     string `csv:"bill_of_lading_date"`
	EquipmentWeight      int64  `csv:"equipment_weight"`
	TareWeight           int64  `csv:"tare_weight"`
	AllowableWeight      bool   `csv:"allowable_weight"`
	DunnageWeight        bool   `csv:"dunnage_weight"`
	EquipmentWeightCode  string `csv:"equipment_weight_code"`
	CommodityCode        int64  `csv:"commodity_code"`
	CommodityDescription string `csv:"commodity_description"`
	OriginID             string `csv:"origin_id"`
	DestinationID        bool   `csv:"destination_id"`
	Routes               string `csv:"routes"`
	Parties              string `csv:"parties"`
}

type Event struct {
	ID                    int64  `csv:"id"`
	EquipmentID           string `csv:"equipment_id"`
	SightingDate          string `csv:"sighting_date"`
	SightingEventCode     int64  `csv:"sighting_event_code"`
	ReportingRailroadSCAC string `csv:"reporting_railroad_scac"`
	PostingDate           string `csv:"posting_date"`
	FromMarkID            string `csv:"from_mark_id"`
	LoadEmptyStatus       string `csv:"load_empty_status"`
	SightingClaimCode     string `csv:"sighting_claim_code"`
	SightingEventCodeText string `csv:"sighting_event_code_text"`
	TrainID               string `csv:"train_id"`
	TrainAlphaCode        string `csv:"train_alpha_code"`
	LocationID            int64  `csv:"location_id"`
	WaybillID             int64  `csv:"waybill_id"`
}
