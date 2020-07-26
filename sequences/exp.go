package sequences

import (
	"math"
)

type Sequence interface {
	Get(i int) float64
	Length() int
}

type Fixed struct {
	data []float64
}

func (f Fixed) Get(i int) float64 {
	return f.data[i]

}

func (f Fixed) Length() int {
	return len(f.data)
}

type Exp struct {
	k      float64
	length int
}

func (e Exp) Get(i int) (v float64) {
	return math.Exp(float64(i) * e.k)
}
func (e Exp) Length() int {
	return e.length
}
