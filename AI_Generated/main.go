package main

import (
	"fmt"
	"log"
	"time"

	"github.com/montanaflynn/stats"
)

// makeCoordinates creates a slice of stats.Coordinate from x and y values.
func makeCoordinates(x, y []float64) ([]stats.Coordinate, error) {
	if len(x) != len(y) {
		return nil, fmt.Errorf("mismatched lengths: x=%d, y=%d", len(x), len(y))
	}
	var coords []stats.Coordinate
	for i := range x {
		coords = append(coords, stats.Coordinate{X: x[i], Y: y[i]})
	}
	return coords, nil
}

// calculateRegression computes the slope and intercept of a linear regression line.
func calculateRegression(coords []stats.Coordinate) (slope, intercept float64, err error) {
	regLine, err := stats.LinearRegression(coords)
	if err != nil {
		return 0, 0, err
	}
	// Use the first and last points of the regression line to compute slope
	slope = (regLine[len(regLine)-1].Y - regLine[0].Y) / (regLine[len(regLine)-1].X - regLine[0].X)
	intercept = regLine[0].Y - slope*regLine[0].X
	return slope, intercept, nil
}

// anscombeData returns the four datasets of the Anscombe Quartet.
func anscombeData() []struct{ x, y []float64 } {
	return []struct {
		x, y []float64
	}{
		{
			x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			y: []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
		},
		{
			x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			y: []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
		},
		{
			x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			y: []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
		},
		{
			x: []float64{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 19},
			y: []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 5.56, 7.91, 6.89, 12.50},
		},
	}
}

func main() {
	start := time.Now() // Start timing

	datasets := anscombeData()

	for i, d := range datasets {
		coords, err := makeCoordinates(d.x, d.y)
		if err != nil {
			log.Fatalf("Dataset %d coordinate error: %v", i+1, err)
		}

		slope, intercept, err := calculateRegression(coords)
		if err != nil {
			log.Fatalf("Dataset %d regression error: %v", i+1, err)
		}

		fmt.Printf("Dataset %d: Slope = %.2f, Intercept = %.2f\n", i+1, slope, intercept)
	}

	duration := time.Since(start) // End timing
	fmt.Printf("Execution time: %v\n", duration)
}
