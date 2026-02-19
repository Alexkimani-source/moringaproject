# 1. Title & Objective
## Prompt-Powered Kickstart: Building a Location-to-Coordinates CLI with Go

**Overview**

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

## AI Prompt Journal (Structured Go Learning Journey)

This journal documents the step-by-step learning process of building the Location-to-Coordinates CLI using Go and the [OpenStreetMap Nominatim API](https://www.openstreetmap.org/#map=6/0.17/37.90).
It shows how AI was used as a tutor to move from basic concepts to a working real-world tool.
**Phase 1: Conceptual Understanding & Setup**

This phase focused on building foundational knowledge about Go and APIs before implementation.

**Prompt 1 — Understanding Go for CLI Development**

**Prompt:**
“Explain how Go works for building command-line tools and how it handles HTTP requests and JSON compared to other languages.”

**AI Helpfulness:**
This clarified Go’s strengths such as simplicity, compiled binaries, and built-in concurrency.
It also explained the ``` net/http ```and ```encoding/json``` packages, which became essential for the project.

## Prompt 2 — Understanding Geocoding APIs 

**Prompt:**
“What is a geocoding API and how does the OpenStreetMap Nominatim API convert place names into coordinates?”

**AI Helpfulness:**
Helped establish the mental model of forward geocoding (name → coordinates) and how query parameters work in REST APIs.

## Prompt 3 — Environment Setup 
**Prompt:**
“Guide me step-by-step to install Go, create a project folder, and run my first Go program.”

**AI Helpfulness:**
Provided installation steps, explained GOPATH, and ensured the environment was correctly configured before coding.

## Phase 2: Building the CLI MVP**

This phase focused on creating the first working version of the toolkit.

**Prompt 4 — Making the First API Request**

**Prompt:**
“Show me how to send a GET request in Go to an API endpoint and print the response.”

**AI Helpfulness:**
Provided the base HTTP request structure and error handling pattern used throughout the project.

**Prompt 5 — Parsing JSON Response**

**Prompt:**
“Help me create a Go struct to parse latitude and longitude from a JSON response.”

**AI Helpfulness:**
Explained struct tags and decoding JSON into Go types, enabling structured output.

**Prompt 6 — Handling User Input**

**Prompt:**
“How do I read user input from the terminal and pass it as a query parameter in Go?”

**AI Helpfulness:**
This enabled the transition from a static request to an interactive CLI.

## Phase 3: Error Handling & UX Improvements

This phase improved reliability and user experience.

**Prompt 7 — Handling ‘Location Not Found’**

**Prompt:**
“How can I check if the API returns no results and display a user-friendly message?”

**AI Helpfulness:**
Helped implement conditional logic to detect empty responses and print
 “Area not found”.

**Prompt 8 — Debugging API Issues**

**Prompt:**
“My Go program compiles but doesn’t return data. How can I debug API requests step-by-step?”

**AI Helpfulness:**
Introduced checking status codes, printing raw responses, and validating URLs.

## Phase 4: Reflection & Learning Outcomes

This phase captures insights gained from the project.

**Key Skills Learned**

Making HTTP requests in Go
Parsing JSON into structs
Working with real-world APIs
Building interactive CLI tools
Debugging network errors

**Conceptual Takeaways**

APIs are contracts between services
Error handling is essential for reliability
Go’s simplicity makes it ideal for small tools

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

Go official documentation  [ GO ](go.dev/doc/)

OpenStreetMap Nominatim API [OSM](https://www.openstreetmap.org/#map=6/0.17/37.90)

JSON parsing tutorial [w3schools]( https://www.w3schools.com/js/js_json_parse.asp) 
