package measurer

import (
	"github.com/peter9207/anna/sequences"
)

func MeanSquared(seq1, seq2 sequences.Sequence) (diff float64) {

	sum := float64(0)
	for i := 0; i < seq1.Length(); i++ {
		d := seq1.Get(i) - seq2.Get(i)
		sum += d * d
	}

	diff = sum / float64(seq1.Length())
	return
}
