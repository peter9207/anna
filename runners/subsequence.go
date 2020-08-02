package runners

import "github.com/peter9207/anna/measurer"
import "github.com/peter9207/anna/sequences"
import "sort"
import "fmt"

type Subsequence struct {
	Data []float64
}

type Result struct {
	Start int
	End   int
	Score float64
}

func (r Result) String() string {
	return fmt.Sprintf("Start: %v, End: %v, Score: %v", r.Start, r.End, r.Score)
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
			results = append(results, Result{
				Start: start,
				End:   end,
				Score: val,
			})
		}
	}
	sort.Slice(results,
		func(i, j int) bool {
			return results[i].Score < results[j].Score
		})
	return

}

func calculateEstimator(k float64, offset int64, sequence []float64) float64 {

	exp := sequences.Exp{K: k, Offset: offset, Len: len(sequence)}
	f := sequences.Fixed{Data: sequence}

	return measurer.MeanSquared(exp, f)
}
