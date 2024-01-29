package handler

import (
	"encoding/json"
	"net/http"
)

type FlightList struct {
	Flights []Flight `json:"flights"`
}

// Flight represents a flight path with a source and destination.
type Flight struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

// findStartingAirport returns the starting airport of the journey.
func findStartingAirport(flights []Flight) string {
	airportMap := make(map[string]int)
	for _, flight := range flights {
		airportMap[flight.Source]--
		airportMap[flight.Destination]++
	}
	for airport, count := range airportMap {
		if count == -1 {
			return airport
		}
	}
	return ""
}

// SortFlights sorts the flights and returns the start and end airports.
func SortFlights(flights []Flight) []string {
	if len(flights) <= 1 {
		if len(flights) == 1 {
			return []string{flights[0].Source, flights[0].Destination}
		}
		return []string{}
	}

	start := findStartingAirport(flights)
	if start == "" {
		return []string{}
	}

	flightMap := make(map[string]string)
	for _, flight := range flights {
		flightMap[flight.Source] = flight.Destination
	}

	end := start

	for {
		next, exists := flightMap[end]
		if !exists {
			break
		}
		end = next
	}

	return []string{start, end}
}

// CalculateHandler handles the http request
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	var flightList FlightList
	if err := json.NewDecoder(r.Body).Decode(&flightList); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(flightList.Flights) == 0 {
		http.Error(w, "At least one flight entry is required", http.StatusBadRequest)
		return
	}

	sortedFlights := SortFlights(flightList.Flights)
	response, err := json.Marshal(sortedFlights)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
