package comdatatype

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"released"`
	Color  bool     `json:"color,omitempty"`
	Actors []string `json:"actors"`
}

func TestJson(t *testing.T) {
	movie := Movie{
		Title:  "Casablanca",
		Year:   1942,
		Color:  false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	}
	movieStr, err := json.Marshal(movie)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", movieStr)
}

func TestJsonUnmarshal(t *testing.T) {
	movieStr := `{"title":"Casablanca","released":1942,"actors":["Humphrey Bogart","Ingrid Bergman"]}`
	var movie Movie
	err := json.Unmarshal([]byte(movieStr), &movie)
	if err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%v\n", movie)
}
