package retry

type Handler func() error

type Retry interface {
	Retry(handler Handler) error
}
