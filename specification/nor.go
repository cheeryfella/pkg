package specification

type NOr[T any] struct {
	Spec[T]
	a Specification[T]
	b Specification[T]
}

func (s NOr[T]) IsSatisfiedBy(t T) bool {
	return !(s.a.IsSatisfiedBy(t) || s.b.IsSatisfiedBy(t))
}
