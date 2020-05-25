package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"math"
)

func findPOI(values []float64) (indexes []int64) {

	windowSize := 9

	threshold := float64(1000)

	// var start, end int64

	// start = 0
	// end = int64(windowSize)

	for i := 0; i < len(values)-windowSize; i++ {

		_, v1 := stat.MeanVariance(values[i:i+4], nil)
		_, v2 := stat.MeanVariance(values[i+4:i+windowSize], nil)

		diff := v1 - v2

		if math.Abs(diff) > threshold {
			fmt.Printf("first Half: %v second half: %v, mid: %v\n", v1, v2, values[i])
			indexes = append(indexes, int64(i+4))
		}

	}
	return
}
