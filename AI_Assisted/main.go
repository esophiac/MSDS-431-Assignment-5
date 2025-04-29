package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

// calculate the mean of the variables in a stats.Coordinate object
// the first return is X and the second return is y
func varMeans(inputData []stats.Coordinate) (xMean float64, yMean float64) {
	var xSum, ySum float64

	// Calculate the sum of x and y values in a single loop
	for i := 0; i < len(inputData); i++ {
		xSum += inputData[i].X
		ySum += inputData[i].Y
	}

	// Calculate the means
	xMean = xSum / float64(len(inputData))
	yMean = ySum / float64(len(inputData))

	return xMean, yMean
}

// calculates the gradient of a stats.Coordinate object
func gradient(inputData []stats.Coordinate) (gradient float64) {
	// Calculate gradient using Sxy / Sxx
	xMean, yMean := varMeans(inputData)
	var numerator, denominator float64

	for _, point := range inputData {
		xDiff := point.X - xMean
		yDiff := point.Y - yMean
		numerator += xDiff * yDiff
		denominator += xDiff * xDiff
	}

	gradient = numerator / denominator

	return gradient
}

// calculates the intersection of a stats.Coordinate object
func yIntercept(inputData []stats.Coordinate) (intercept float64) {
	// Precompute gradient and means
	lineGradient := gradient(inputData)
	xMean, yMean := varMeans(inputData)

	// Calculate the y-intercept
	intercept = yMean - lineGradient*xMean

	return intercept
}

// Make it easier to print linea regression coefficients
// this function takes a dataset and a label and prints the gradient and intercept
func processDataset(dataset []stats.Coordinate, label string) {
	gradientValue := gradient(dataset)
	yInterceptValue := yIntercept(dataset)
	fmt.Printf("%s - Gradient: %.2f, Intercept: %.2f\n", label, gradientValue, yInterceptValue)
}

func main() {

	// Define the points in the Anscombe Quartet
	num1 := []stats.Coordinate{{10, 8.04}, {8, 6.95}, {13, 7.58}, {9, 8.81}, {11, 8.33}, {14, 9.96}, {6, 7.24}, {4, 4.26}, {12, 10.84}, {7, 4.82}, {5, 5.68}}
	num2 := []stats.Coordinate{{10, 9.14}, {8, 8.14}, {13, 8.74}, {9, 8.77}, {11, 9.26}, {14, 8.1}, {6, 6.13}, {4, 3.1}, {12, 9.13}, {7, 7.26}, {5, 4.74}}
	num3 := []stats.Coordinate{{10, 7.46}, {8, 6.77}, {13, 12.74}, {9, 7.11}, {11, 7.81}, {14, 8.84}, {6, 6.08}, {4, 5.39}, {12, 8.15}, {7, 6.42}, {5, 5.73}}
	num4 := []stats.Coordinate{{8, 6.58}, {8, 5.76}, {8, 7.71}, {8, 8.84}, {8, 8.47}, {8, 7.04}, {8, 5.25}, {19, 12.5}, {8, 5.56}, {8, 7.91}, {8, 6.89}}

	// Process each dataset
	processDataset(num1, "Dataset 1")
	processDataset(num2, "Dataset 2")
	processDataset(num3, "Dataset 3")
	processDataset(num4, "Dataset 4")

}
