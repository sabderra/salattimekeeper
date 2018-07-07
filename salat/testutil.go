package salat

import (
	"fmt"
	"math"
	"testing"
)

func assertEquals(test *testing.T, actual string, expected string, message string) {
	if actual != expected {
		errMsg := fmt.Sprintf("%s:, expected %s, got %s", message, expected, actual)
		test.Error(errMsg)
	}
}

func assertFloatEquals(test *testing.T, actual float64, expected float64, message string) {
	if actual-expected != 0 {
		errMsg := fmt.Sprintf("%s:, expected %.16f, got %.16f", message, expected, actual)
		test.Error(errMsg)
	}
}

func assertClose(test *testing.T, actual float64, expected float64, precision int, message string) {
	diff := toFixed(actual, precision) - toFixed(expected, precision)
	if diff != 0 {
		errMsg := fmt.Sprintf("%s:, expected %f, got %f", message, expected, actual)
		test.Error(errMsg)
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
