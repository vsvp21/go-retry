package retry

import "math/rand"

type Jitter interface {
	Calc(delay float64) float64
}

func NewFullJitter(minDelay float64) *FullJitter {
	return &FullJitter{
		minDelay: minDelay,
	}
}

type FullJitter struct {
	minDelay float64
}

func (j *FullJitter) Calc(delay float64) float64 {
	return j.minDelay + rand.Float64()*(delay-j.minDelay)
}
