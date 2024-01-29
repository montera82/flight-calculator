package handler_test

import (
	"flights-calculator/pkg/handler"
	"reflect"
	"testing"
)

func TestSortFlights(t *testing.T) {
	testCases := []struct {
		name     string
		input    []handler.Flight
		expected []string
	}{
		{
			name: "It should calculate correctly",
			input: []handler.Flight{
				{Source: "IND", Destination: "EWR"},
				{Source: "SFO", Destination: "ATL"},
				{Source: "GSO", Destination: "IND"},
				{Source: "ATL", Destination: "GSO"},
			},
			expected: []string{"SFO", "EWR"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := handler.SortFlights(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Test %s failed. Expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
