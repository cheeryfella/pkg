package specification

type Not[T any] struct {
	Spec[T]
	a Specification[T]
}

func (s Not[T]) IsSatisfiedBy(t T) bool {
	return !s.a.IsSatisfiedBy(t)
}
