package main

import (
	"fmt"
	"math"
	"time"

	"github.com/montanaflynn/stats"
)

type SummaryStats struct {
	Intercept              float64
	Slope                  float64
	CorrelationCoefficient float64
	RSquared               float64
	StandardError          float64
}

// Given x, y, and the mean of x and y values, get the difference for each point
// Returns: Sum of squared differences
func CalculateDiffs(x, y []float64, mean_X, mean_Y float64) (diff_XY, diff_XX, diff_YY float64) {
	for i, xi := range x {
		diff_X := xi - mean_X
		diff_Y := y[i] - mean_Y
		diff_XY += diff_X * diff_Y
		diff_XX += diff_X * diff_X
		diff_YY += diff_Y * diff_Y
	}
	return
}

// Given x, y, and the slope and intercept of a best-fit line, get residual error
// Returns: squared residuals
func CalculateResiduals(x, y []float64, slope, intercept float64) (squared_residuals float64) {
	for i, xi := range x {
		yi := slope*xi + intercept
		residual := y[i] - yi
		squared_residuals += residual * residual
	}
	return
}

// Given x, y, bring the two above functions together & produce a linear model summary
// Returns: SummaryStats
func LmSummary(x, y []float64) SummaryStats {
	n := len(x)
	mean_X, _ := stats.Mean(x)
	mean_Y, _ := stats.Mean(y)

	diff_XY, diff_XX, diff_YY := CalculateDiffs(x, y, mean_X, mean_Y)

	slope := diff_XY / diff_XX
	intercept := mean_Y - slope*mean_X

	r := diff_XY / math.Sqrt(diff_XX*diff_YY)
	rSq := r * r

	squared_residuals := CalculateResiduals(x, y, slope, intercept)
	standard_error := math.Sqrt(squared_residuals / float64(n-2))

	return SummaryStats{
		Intercept:              intercept,
		Slope:                  slope,
		CorrelationCoefficient: r,
		RSquared:               rSq,
		StandardError:          standard_error,
	}
}

func main() {
	start_time := time.Now()
	// the slice of slices makes iterating through each anscombe quadrant easier
	x := [][]float64{
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
	}

	y := [][]float64{
		{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
		{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
		{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
		{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
	}
	for i := range x {
		fmt.Printf("Quadrant %d\n", i+1)
		stats := LmSummary(x[i], y[i])
		fmt.Printf("Intercept: %v\n", stats.Intercept)
		fmt.Printf("Slope: %v\n", stats.Slope)
		fmt.Printf("Correlation Coefficient: %v\n", stats.CorrelationCoefficient)
		fmt.Printf("RSquared Value: %v\n", stats.RSquared)
		fmt.Printf("Standard Error of Estimate: %v\n", stats.StandardError)
		fmt.Println("----------------------------------------------------")
	}
	elapsed := time.Since(start_time)
	fmt.Printf("Took: %s", elapsed)

}
