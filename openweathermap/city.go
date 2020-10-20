package openweathermap

// City ...
type City struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Country    string     `json:"country"`
	Coord      Coordinate `json:"coord"`
	Population int32      `json:"population"`
}

// Coordinate ...
type Coordinate struct {
	Longitude float32 `json:"lon"`
	Latitude  float32 `json:"lat"`
}
