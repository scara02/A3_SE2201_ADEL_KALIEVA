package tests

import (
	"testing"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Movie struct {
	Title    string   `json:"title"`
	Year     int      `json:"year"`
	Runtime  string   `json:"runtime"`
	Genres   []string `json:"genres"`
}

// First I added permission "movies:write" for my user
func TestCreateMovie(t *testing.T) {
	moviePayload := Movie{
		Title:   "Twilight",
		Year:    2008,
		Runtime: "121 mins",
		Genres:  []string{"romance"},
	}

	payloadBytes, err := json.Marshal(moviePayload)
	if err != nil {
		log.Fatalf("Error marshaling movie payload: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "WVO66BUNR3GS2UJBF23MRS6SYU" // authentication_token for my user
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer res.Body.Close()

	printResponseBody(res)
}

func TestCreateMovie_InvalidTitle(t *testing.T) {
	moviePayload := Movie{
		Title:   "",
		Year:    2008,
		Runtime: "121 mins",
		Genres:  []string{"romance"},
	}

	payloadBytes, err := json.Marshal(moviePayload)
	if err != nil {
		log.Fatalf("Error marshaling movie payload: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "WVO66BUNR3GS2UJBF23MRS6SYU"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer res.Body.Close()

	printResponseBody(res)
}

func TestCreateMovie_InvalidYear(t *testing.T) {
	moviePayload := Movie{
		Title:   "Twilight",
		Year:    2050,
		Runtime: "121 mins", 
		Genres:  []string{"romance"},
	}

	payloadBytes, err := json.Marshal(moviePayload)
	if err != nil {
		log.Fatalf("Error marshaling movie payload: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/movies", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "WVO66BUNR3GS2UJBF23MRS6SYU"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer res.Body.Close()

	printResponseBody(res)
}

func TestDeleteMovie(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:4000/v1/movies/1", nil)
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	bearerToken := "WVO66BUNR3GS2UJBF23MRS6SYU"
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer res.Body.Close()
	
	printResponseBody(res)
}