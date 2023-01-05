package model

type EmissionsMessage struct {
	Tail_number  string  // The international aircraft registration. A unique code assigned to the aircraft.
	Engine_count int     // The number of engines on the aircraft.
	Engine_name  string  // The engine name.
	Latitude     float64 // The Latitude in degrees.
	Longitude    float64 // The Longitude in degrees.
	Altitude     float64 // The Altitude in degrees.
	Temperature  float64 // The Temperature in degrees Fahrenheit.
}
