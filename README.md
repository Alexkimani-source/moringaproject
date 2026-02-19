# moringaproject
A Go CLI tool that fetches coordinates for a location using OpenStreetMap Nominatim API.
# Prompt-Powered Kickstart: Building a Location-to-Coordinates CLI with Go

# Overview
This project is a command-line application built with Go that converts a location name into geographic coordinates (latitude and longitude) using the OpenStreetMap Nominatim API.
## Technology chosen : GO
**Why I chose it:**
I chose Go because it is efficient, beginner-friendly, and widely used for backend services and APIs. I wanted to learn how to make HTTP requests, handle JSON data, and work with real-world API responses.
**End goal:**
To build a command-line application that fetches latitude and longitude for a given location using the OpenStreetMap Nominatim API, demonstrating real-world API integration.

## Quick Summary of the Technology
**What is Go?**
Go is an open-source programming language developed by Google for building fast and reliable software.

**Where it’s used:**
1. Backend systems

2. Cloud tools

3. APIs and microservices

**Real-world example:**
Go is used in cloud infrastructure tools, API services, and web backends.

## System Requirements
OS: Windows / Linux / Mac

Editor: Visual Studio Code

Go installed

Internet connection (for API request)

**Installation & Setup Instructions**

1. Install Go

Download and install Go from the official website

2. Verify Installation
   
   ``` go version```

3. Create Project Folder
   ```mkdir location-cli ```
   ```cd location-cli```
   
4. Create File
  ``` main.go```


## Working Example

**What the program does**:
The program sends a request to the OpenStreetMap Nominatim API, retrieves latitude and longitude for a user-provided location, and prints it in the terminal. If no result is found, it reports “Area not found.”

```
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
```
**Run Command**
``` go run main.go "Nairobi, Kenya" ```

## Expected Output Example

**Valid location:**
```
Location: Nairobi, Nairobi County, Kenya
Latitude: -1.2920659
Longitude: 36.8219462
```
**Invalid location:**
```
Area not found: RandomPlaceThatDoesNotExist
```
## AI Prompt Journal

**Prompt 1**
Prompt: “Explain how to make an HTTP request in Go”
**How AI helped:** Provided the basic structure using net/http and explained response handling.

**Prompt 2**
Prompt: “Show me how to parse JSON in Go”
How AI helped: Helped define a struct and decode JSON response.

**Prompt 3**
Prompt: “How to handle no results returned by an API in Go”
**How AI helped:** Explained checking array length and returning nil if no results are found.

## Common Issues & Fixes
**1. API request fails**
   The program couldn’t fetch data from the API (e.g., internet down, API server unreachable).
   
**Fix:**
   i. Check that the internet connection is working.

   ii. Ensure the API URL is correct.

   iii. Implement error handling to prevent the program from crashing if the API does not respond.

**2. JSON fields empty**
Struct fields are not exported.

**Fix**: Capitalize struct fields.

**3. go command not recognized**
Couldn’t run ```go run main.go``` in the terminal.

**Fix**: Added Go bin directory (```C:\Program Files\Go\bin```) to the system PATH and reopened terminal.

**Directory structure when running the project:**
```
C:\Users\<username>\Documents\location-cli\
│
├─ main.go
├─ README.md
└─ go.mod
```
**AI Agents Utilised**
Claude , 
ChatGPT, 
Gemini CLI

**References**

Go official documentation

OpenStreetMap Nominatim API

JSON parsing tutorial
