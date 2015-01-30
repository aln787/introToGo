//Visit in browser http://api.openweathermap.org/data/2.5/weather?q=spotsylvania
//View in Json editor https://www.jsoneditoronline.org/
//View chorme developer tools network request

//Get the json structure using gojson
// `cd /Users/xvo202/go/src/github.com/ChimeraCoder/gojson`
// `curl -s http://api.openweathermap.org/data/2.5/weather?q=spotsylvania | gojson`
// gojson repo `https://github.com/ChimeraCoder/gojson.git`

package main

type Foo struct {
	Base   string `json:"base"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Cod   int `json:"cod"`
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Dt   int `json:"dt"`
	ID   int `json:"id"`
	Main struct {
		Humidity int     `json:"humidity"`
		Pressure float64 `json:"pressure"`
		Temp     float64 `json:"temp"`
		TempMax  float64 `json:"temp_max"`
		TempMin  float64 `json:"temp_min"`
	} `json:"main"`
	Name string `json:"name"`
	Sys  struct {
		Country string  `json:"country"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
		Type    int     `json:"type"`
	} `json:"sys"`
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
		ID          int    `json:"id"`
		Main        string `json:"main"`
	} `json:"weather"`
	Wind struct {
		Deg   float64 `json:"deg"`
		Speed float64 `json:"speed"`
	} `json:"wind"`
}
