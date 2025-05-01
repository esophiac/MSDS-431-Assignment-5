package main

import (
	"testing"

	"github.com/montanaflynn/stats"
)

// TestMakeCoordinates checks that makeCoordinates correctly converts x and y to coordinates
func TestMakeCoordinates(t *testing.T) {
	x := []float64{1, 2, 3}
	y := []float64{4, 5, 6}
	coords, err := makeCoordinates(x, y)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(coords) != 3 {
		t.Errorf("Expected 3 coordinates, got %d", len(coords))
	}
	if coords[0].X != 1 || coords[0].Y != 4 {
		t.Errorf("Expected (1,4), got (%.2f, %.2f)", coords[0].X, coords[0].Y)
	}
}

// TestMakeCoordinates_Mismatch checks for mismatched x and y lengths
func TestMakeCoordinates_Mismatch(t *testing.T) {
	x := []float64{1, 2}
	y := []float64{4}
	_, err := makeCoordinates(x, y)
	if err == nil {
		t.Error("Expected error for mismatched lengths, got nil")
	}
}

// TestCalculateRegression checks slope/intercept from simple coordinates
func TestCalculateRegression(t *testing.T) {
	coords := []stats.Coordinate{
		{X: 1, Y: 2},
		{X: 2, Y: 4},
		{X: 3, Y: 6},
	}
	slope, intercept, err := calculateRegression(coords)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expectedSlope := 2.0
	expectedIntercept := 0.0

	if slope != expectedSlope {
		t.Errorf("Expected slope %.2f, got %.2f", expectedSlope, slope)
	}
	if intercept != expectedIntercept {
		t.Errorf("Expected intercept %.2f, got %.2f", expectedIntercept, intercept)
	}
}
