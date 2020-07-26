package sequences

import (
	"math"
)

type Sequence interface {
	Get(i int) float64
	Length() int
}

type Fixed struct {
	Data []float64
}

func (f Fixed) Get(i int) float64 {
	return f.Data[i]

}

func (f Fixed) Length() int {
	return len(f.Data)
}

type Exp struct {
	K      float64
	Len    int
	Offset int64
}

func (e Exp) Get(i int) (v float64) {
	return math.Exp(float64(i)*e.K) + float64(e.Offset)
}
func (e Exp) Length() int {
	return e.Len
}
