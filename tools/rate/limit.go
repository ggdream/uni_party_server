package rate

type Limiter interface {
	Take() error
}
