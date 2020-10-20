# OpenWeatherMap API Client for Go

![Go](https://github.com/otiai10/openweathermap/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/otiai10/openweathermap/branch/main/graph/badge.svg?token=5TyeTM6vgn)](https://codecov.io/gh/otiai10/openweathermap)

```go
package main

import (
  "github.com/otiai10/openweathermap"
)
func main() {
  client := openweathermap.New(API_KEY)
  res, _ := client.ByCityName("Tokyo")
  fmt.Println(res.Forecasts[0].Weather[0].Main)
  // Clouds
}
```
