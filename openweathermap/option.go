package openweathermap

import (
	"fmt"
	"net/url"
)

// Option ...
type Option struct {
	Units string
	Count int
}

// DefaultOption ...
var DefaultOption = Option{
	Units: "metric",
	Count: 10,
}

// Query ...
func (opt *Option) Query() url.Values {
	v := url.Values{
		"units": {opt.Units},
		"count": {fmt.Sprintf("%d", opt.Count)},
	}
	return v
}
