# Go Weather CLI 🌦️

A lightning-fast, command-line weather application built entirely in Go. 

This project was developed to demonstrate proficiency in system-level Go programming, API consumption, strict type mapping, and JSON unmarshaling using only the Go Standard Library.

## Features
* **Instant Weather Data:** Fetches real-time temperature, humidity, and condition descriptions for any global city.
* **Zero Dependencies:** Built utilizing only the Go standard library (`net/http`, `encoding/json`, `os`)—no external frameworks required.
* **Error Handling:** Gracefully handles invalid city names, missing arguments, and API connection failures.

## Tech Stack
* **Language:** Go (Golang)
* **API:** OpenWeatherMap REST API
* **Data Format:** JSON

## Prerequisites
To run this project, you will need:
* [Go](https://go.dev/dl/) installed on your machine.
* A free API key from [OpenWeatherMap](https://openweathermap.org/api).

## Installation & Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/ssteve221/weather-cli.git
   cd weather-cli
   ```
2. Open main.go and replace your api here with and API from openweathermap
3. Build the executable (optional but recommended for speed):
```bash
go build -o weather main.go
```

Usage
Run the program by passing a city name as an argument. If the city name contains spaces, wrap it in quotes.

```bash
# Using go run
go run main.go London
```
```bash
# Using the built executable
./weather "New York"
```
Example Output
💻$_:
🌍 Weather in London:
🌡️  Temperature: 15.2°C
💧 Humidity: 72%
☁️  Conditions: scattered clouds

What I Learned:
Building this project reinforced my understanding of taking CLI arguments in Go, establishing safe HTTP client connections, and creating precise structs to unmarshal external JSON payloads into memory-safe Go data structures.
