package main

import (
	"testing"

	"github.com/montanaflynn/stats"
	"github.com/stretchr/testify/assert"
)

func TestCalculateDiffs(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	y := []float64{22, 33, 44, 55, 66, 77, 88, 99, 88, 77, 66}
	meanX, _ := stats.Mean(x)
	meanY, _ := stats.Mean(y)
	diffXY, diffXX, diffYY := CalculateDiffs(x, y, meanX, meanY)
	assert.Equal(t, 638.0, diffXY, "diffXY was incorrect")
	assert.Equal(t, 110.0, diffXX, "diffXX was incorrect")
	assert.Equal(t, 5918.0, diffYY, "diffYY was incorrect")
}

func TestCalculateResiduals(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	y := []float64{22, 33, 44, 55, 66, 77, 88, 99, 88, 77, 66}
	slope := 5.8
	intercept := 30.2
	squared_residuals := CalculateResiduals(x, y, slope, intercept)
	assert.InDelta(t, 2217.6, squared_residuals, 0.001, "squared_residuals was incorrect")
}

func TestLmSummary(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	y := []float64{22, 33, 44, 55, 66, 77, 88, 99, 88, 77, 66}
	stats := LmSummary(x, y)
	assert.InDelta(t, 30.2, stats.Intercept, .001, "Intercept was incorrect")
	assert.Equal(t, 5.8, stats.Slope, "Slope was incorrect")
	assert.InDelta(t, 0.791, stats.CorrelationCoefficient, .001, "CorrelationCoefficient was incorrect")
	assert.InDelta(t, 0.625, stats.RSquared, .001, "RSquared was incorrect")
	assert.InDelta(t, 15.697, stats.StandardError, .001, "Standard Error was incorrect")
}
