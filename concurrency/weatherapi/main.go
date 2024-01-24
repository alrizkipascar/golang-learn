package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const apiKey = "SECRETAPI KEK W"

func fetchWeather(city string) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
		return data
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather data for %s: %s\n", city, err)
		return data
	}
	return data
}

func fetchWeatherConcurrency(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	defer wg.Done()
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
		return data
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather data for %s: %s\n", city, err)
		return data
	}

	ch <- fmt.Sprintf("This is the %s", city)
	return data
}

func main() {
	startNow := time.Now()

	cities := []string{"Toronto", "Paris", "Tokyo"}
	ch := make(chan string)
	var wg sync.WaitGroup
	for _, city := range cities {
		data := fetchWeather(city)
		fmt.Println("This is the data", data)
		fmt.Println("Time operation:", time.Since(startNow))

	}

	for _, city := range cities {
		wg.Add(1)
		go fetchWeatherConcurrency(city, ch, &wg)
		// fmt.Println("This is the data", data)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
		fmt.Println("Time operation:", time.Since(startNow))

	}

}
