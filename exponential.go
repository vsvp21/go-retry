package retry

import (
	"math"
	"time"
)

func NewExponentialBackoffWithFullJitter(maxInterval float64, minDelay float64) *Exponential {
	return newExponentialBackoffWithJitter(maxInterval, minDelay, NewFullJitter(minDelay))
}

func NewExponentialBackoff(maxInterval float64, minDelay float64) *Exponential {
	return &Exponential{
		maxInterval: maxInterval,
		attempts:    minDelay,
		jitter:      nil,
	}
}

func newExponentialBackoffWithJitter(maxInterval float64, minDelay float64, jitter Jitter) *Exponential {
	return &Exponential{
		maxInterval: maxInterval,
		attempts:    minDelay,
		jitter:      jitter,
	}
}

type Exponential struct {
	maxInterval float64
	attempts    float64
	jitter      Jitter
}

func (e *Exponential) Retry(f Handler) error {
	var err error

	delay := e.attempts
	for delay <= e.maxInterval {
		if err = f(); err != nil {
			time.Sleep(time.Duration(delay) * time.Second)

			delay = math.Min(e.maxInterval, (math.Pow(2, e.attempts)-1)/2)
			if e.jitter != nil {
				delay = e.jitter.Calc(delay)
			}

			e.attempts++
			continue
		}

		break
	}

	return err
}
