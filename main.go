package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// define structs to map exctly json response
type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func main() {
	//read city from terminal argument
	if len(os.Args) < 2 {
		fmt.Println("Error: please provide a city name")
		fmt.Println("Usage:go run main.go <city>")
		os.Exit(1)
	}
	city := os.Args[1]
	apikey := "PLACE API FROM SITE HERE"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apikey)

	//make http request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to connects to Api %v\n", err)
		os.Exit(1)

	}
	defer resp.Body.Close() //close connection when function ends

	//read raw response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response:%v\n", err)
		os.Exit(1)
	}

	//catch invalid cities pr bad API keys
	if resp.StatusCode != 200 {
		fmt.Printf("API ERROR (status %d): check your city name or API key\n", resp.StatusCode)
		os.Exit(1)
	}

	//parse json into Go struct
	var weather WeatherData
	if err := json.Unmarshal(body, &weather); err != nil {
		fmt.Printf("Failed to parse JSON: %v\n", err)
		os.Exit(1)
	}

	//print formatted output
	fmt.Printf("\n🌍 Weather in %s:\n", weather.Name)
	fmt.Printf("🌡️  Temperature: %.1f°C\n", weather.Main.Temp)
	fmt.Printf("💧 Humidity: %d%%\n", weather.Main.Humidity)

	if len(weather.Weather) > 0 {
		fmt.Printf("☁️  Conditions: %s\n\n", weather.Weather[0].Description)
	}
}
