package sequences

type Linear struct {
	k      float64
	offset float64
}

func (l Linear) get(v int64) float64 {
	return l.k*float64(v) + l.offset
}
