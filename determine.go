package main

import "github.com/spf13/cobra"

import "github.com/peter9207/anna/loaders"
import "github.com/peter9207/anna/measurer"
import "github.com/peter9207/anna/sequences"

import "strconv"
import "log"
import "fmt"
import "math"

var expCoeff = &cobra.Command{
	Use: "exp <filename>",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			cmd.Help()
			return
		}

		filename := args[0]

		data := []float64{}

		loaders.FromCSV(filename, func(input []string) {
			f, err := strconv.ParseFloat(input[1], 64)
			if err != nil {
				log.Println(err)
			}
			data = append(data, f)
		})

		trialK := []float64{}

		for i := 0; i < 100; i++ {
			trialK = append(trialK, float64(i)/1000)
		}

		min := math.MaxFloat64
		minK := float64(0)

		for _, v := range trialK {
			val := calculateEstimator(v, 3, data)
			fmt.Printf("resulting parameters: %v error %v \n", v, val)

			if min > val {
				min = val
				minK = v
			}
		}

		fmt.Printf("min parameters: %v error %v \n", minK, min)

	}}

func calculateEstimator(k float64, offset int64, sequence []float64) float64 {

	exp := sequences.Exp{K: k, Offset: offset, Len: len(sequence)}
	f := sequences.Fixed{Data: sequence}

	return measurer.MeanSquared(exp, f)
}
