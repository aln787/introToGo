//Visit in browser http://api.openweathermap.org/data/2.5/weather?q=spotsylvania
//View in Json editor https://www.jsoneditoronline.org/
//View chorme developer tools network request

//Get the json structure using gojson
// `cd /Users/xvo202/go/src/github.com/ChimeraCoder/gojson`
// `curl -s http://api.openweathermap.org/data/2.5/weather?q=spotsylvania | gojson`
// gojson repo `https://github.com/ChimeraCoder/gojson.git`

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
	//locations := [5]{"spotsylvania", "washington", "chicago", "san francisco", "new york"}
	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=spotsylvania")
	if err != nil {
		panic(err)
		//return weatherData{}, err
	}
	defer res.Body.Close()
	var d weatherData
	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		panic(err)
		//return weatherData{}, err
	}

	w.Write([]byte(fmt.Sprintf("%+v", d)))
	//return d, nil

	//fmt.Printf("%+v\n", d)
	//fmt.Printf("openWeatherMap: %.2f\n", d.Main.Temp)
}

func query(city string) (weatherData, error) {

}
