package openweathermap

import (
	"fmt"
	"time"
)

// ForecastResponse ...
// See https://openweathermap.org/forecast5#JSON
type ForecastResponse struct {
	Code      string     `json:"cod"`
	Message   float32    `json:"message"`
	Count     int        `json:"cnt"`
	City      City       `json:"city"`
	Forecasts []Forecast `json:"list"`
}

// Forecast read only
// See https://openweathermap.org/forecast5#JSON
type Forecast struct {
	Timestamp int64  `json:"dt"`
	Datetime  string `json:"dt_txt"`
	Main      struct {
		Temp       float32 `json:"temp"`
		TempMin    float32 `json:"temp_min"`
		TempMax    float32 `json:"temp_max"`
		Pressure   float32 `json:"pressure"`    // hPa
		SeaLevel   float32 `json:"sea_level"`   // hPa
		GrandLevel float32 `json:"grand_level"` // hPa
		Humidity   int     `json:"humidity"`    // %
	} `json:"main"`
	Weather []Weather `json:"weather"`
	Clouds  struct {
		All int `json:"all"` // %
	} `json:"clouds"`
	Wind struct {
		Speed     float32 `json:"speed"` // m/s
		Direction float32 `json:"deg"`   // Wind direction degree by meteorogically
	} `json:"wind"`
	Rain struct {
		For3h float32 `json:"3h"` // Rain volume for last 3 hours (mm)
	} `json:"rain"`
	Snow struct {
		For3h float32 `json:"3h"` // Snow volume for last 3 hours (mm)
	} `json:"snow"`

	// Extended
	LocalTime time.Time `json:"local_time"`
}

// Weather ...
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// IconURL ...
func (w Weather) IconURL(suffix ...string) string {
	suffix = append(suffix, "")
	return fmt.Sprintf("https://openweathermap.org/img/wn/%s%s.png", w.Icon, suffix[0])
}

// GroupByDate ...
func (res *ForecastResponse) GroupByDate(loc *time.Location) [][]Forecast {
	result := [][]Forecast{}
	tmp := []Forecast{}
	var groupdate int
	for _, f := range res.Forecasts {
		t := time.Unix(f.Timestamp, 0)
		if loc != nil {
			t = t.In(loc)
		}
		f.LocalTime = t     // FIXME: Not sure this SHOULD be here or not
		_, _, d := t.Date() // Date of this forecast
		if d != groupdate {
			if len(tmp) != 0 {
				result = append(result, tmp[:])
				tmp = []Forecast{}
			}
			groupdate = d
		}
		tmp = append(tmp, f)
	}
	if len(tmp) != 0 {
		result = append(result, tmp[:])
	}
	return result
}
