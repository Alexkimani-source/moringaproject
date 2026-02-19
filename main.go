package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Place struct {
	DisplayName string `json:"display_name"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go \"Location Name\"")
		return
	}

	location := strings.Join(os.Args[1:], " ")

	place, err := getCoordinates(location)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if place == nil {
		fmt.Println("Area not found:", location)
		return
	}

	fmt.Println("Location:", place.DisplayName)
	fmt.Println("Latitude:", place.Lat)
	fmt.Println("Longitude:", place.Lon)
}

func getCoordinates(query string) (*Place, error) {
	apiURL := "https://nominatim.openstreetmap.org/search?" +
		"q=" + url.QueryEscape(query) + "&format=json&limit=1"

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	var results []Place
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	if len(results) == 0 {
		return nil, nil
	}

	return &results[0], nil
}
