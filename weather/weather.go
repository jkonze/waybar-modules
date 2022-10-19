package main

import (
	"encoding/json"
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"log"
)

type WaybarOutput struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
	Class   string `json:"class"`
}

var apiKey string
var location string

func GetWeather() {
	fmt.Println(apiKey)
	w, err := owm.NewCurrent("C", "en", apiKey)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = w.CurrentByName(location)
	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := json.Marshal(WaybarOutput{
		Text:    fmt.Sprintf("%.0f", w.Main.Temp),
		Tooltip: "$tooltip",
		Class:   "$class",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(result))
}

func main() {
	GetWeather()
}
