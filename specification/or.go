package specification

type Or[T any] struct {
	Spec[T]
	a Specification[T]
	b Specification[T]
}

func (s Or[T]) IsSatisfiedBy(t T) bool {
	return s.a.IsSatisfiedBy(t) || s.b.IsSatisfiedBy(t)
}
