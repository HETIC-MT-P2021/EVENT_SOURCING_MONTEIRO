package models

type Event struct {
	ID      int    `json:"id"`
	Agg     string `json:"aggregate_id"`
	Handle  string `json:"handle"`
	Payload string `json:"payload"`
}

func GetEvents() []Event {
	var events []Event

	DB.Find(&events)

	return events
}

func CreateEvent(event *Event) {
	// Migrate the schema
	DB.AutoMigrate(&Event{})

	// Create
	DB.Create(event)
}
