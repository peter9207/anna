package runners

import "github.com/peter9207/anna/loaders"
import "github.com/peter9207/anna/measurer"
import "github.com/peter9207/anna/sequences"
import "fmt"
import "sort"

type Subsequence struct {
	Data []float64
}

type Result struct {
	Start int
	End   int
	Score float64
}

func lessResult(i, j Result) bool {
	return i.Score < j.Score
}

func (s Subsequence) Run() (results []Result) {

	trials := []float64{}
	for i := 0; i < 100; i++ {
		trials = append(trials, float64(i)/1000)
	}

	sequenceLength := 50
	for i := 0; i < (len(s.Data) - sequenceLength); i++ {
		start := i
		end := start + sequenceLength
		for _, v := range trials {
			val := calculateEstimator(v, int64(s.Data[i]), s.Data)
			fmt.Printf("resulting parameters: %v error %v \n", v, val)
			results = append(results, Result{
				Start: start,
				End:   end,
				Score: val,
			})
		}
	}
	sort.Slice(results, lessResult)
	return

}

func calculateEstimator(k float64, offset int64, sequence []float64) float64 {

	exp := sequences.Exp{K: k, Offset: offset, Len: len(sequence)}
	f := sequences.Fixed{Data: sequence}

	return measurer.MeanSquared(exp, f)
}
