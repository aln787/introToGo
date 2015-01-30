//Visit in browser http://api.openweathermap.org/data/2.5/weather?q=spotsylvania
//View in Json editor https://www.jsoneditoronline.org/
//View chorme developer tools network request

//Get the json structure using gojson
// `cd /Users/xvo202/go/src/github.com/ChimeraCoder/gojson`
// `curl -s http://api.openweathermap.org/data/2.5/weather?q=spotsylvania | gojson`
// gojson repo `https://github.com/ChimeraCoder/gojson.git`

////Next steps
//difference between Printf and Sprintf
//Accessing values in the data structure
//move query to a function to prep for goroutine
//return values for each stage

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type weatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
	} `json:"sys"`
}

func main() {
	http.HandleFunc("/weather", getWeather)
	http.ListenAndServe(":8080", nil)
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	d := make(chan weatherData, 2) //Setting the buffer to 2
	e := make(chan error)
	go func() {
		data, err := query("spotsylvania")
		if err != nil {
			e <- err
		}
		d <- data //Collecting from channel after getting only one value in the channel
	}()
	data := <-d
	w.Write([]byte(fmt.Sprintf("%+v", data)))
}

func query(city string) (weatherData, error) {
	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city)
	if err != nil {
		return weatherData{}, err
	}
	defer res.Body.Close()
	var d weatherData
	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	return d, nil
}
