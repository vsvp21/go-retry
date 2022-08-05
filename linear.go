package retry

import (
	"time"
)

func NewLinear(delayInSeconds, maxTries int) *Linear {
	return &Linear{
		delayInSeconds: delayInSeconds,
		maxTries:       maxTries,
	}
}

type Linear struct {
	delayInSeconds int
	maxTries       int
}

func (l *Linear) Retry(f Handler) error {
	var err error
	tries := 0

	for tries < l.maxTries {
		tries++
		if err = f(); err != nil && tries != l.maxTries {
			time.Sleep(time.Second * time.Duration(l.delayInSeconds))
			continue
		}

		break
	}

	return err
}
