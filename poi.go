package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"math"
)

type POI struct {
	index int64
	value float64
}

func findPOI(values []float64) (indexes []POI) {

	windowSize := 9
	threshold := float64(1000)

	for i := 0; i < len(values)-windowSize; i++ {

		_, v1 := stat.MeanVariance(values[i:i+4], nil)
		_, v2 := stat.MeanVariance(values[i+4:i+windowSize], nil)

		diff := v1 - v2

		if math.Abs(diff) > threshold {
			fmt.Printf("index: %v,  first Half: %v second half: %v, mid: %v\n", i, v1, v2, values[i])
			poi := POI{
				index: int64(i + 4),
				value: math.Abs(diff),
			}
			indexes = append(indexes, poi)
		}
	}
	return
}
