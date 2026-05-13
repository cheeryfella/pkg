package specification

type Specification[T any] interface {
	// IsSatisfiedBy check if entity is satisfied by the specification.
	IsSatisfiedBy(t T) bool
	// And create a new specification that is the AND operation of the current specification and
	// another specification.
	And(spec Specification[T]) Specification[T]
	// Or create a new specification that is the OR operation of the current specification and
	// another specification.
	Or(spec Specification[T]) Specification[T]
	// Not create a new specification that is the NOT operation of the current specification.
	Not() Specification[T]
	// NAnd create a new specification that is the NOT and NOT operation of the current specification anf
	// another specification
	NAnd(spec Specification[T]) Specification[T]
	// Not create a new specification that is the NOT operation of the current specification.
	NOr(spec Specification[T]) Specification[T]
}

type Spec[T any] struct {
	Specification[T]
}

func (s *Spec[T]) And(other Specification[T]) Specification[T] {
	ns := &And[T]{a: s.Specification, b: other}
	ns.Associate(ns)
	return ns
}

func (s *Spec[T]) NAnd(other Specification[T]) Specification[T] {
	ns := &NAnd[T]{a: s.Specification, b: other}
	ns.Associate(ns)
	return ns
}

func (s *Spec[T]) NOr(other Specification[T]) Specification[T] {
	ns := &NOr[T]{a: s.Specification, b: other}
	ns.Associate(ns)
	return ns
}

func (s *Spec[T]) Not() Specification[T] {
	ns := &Not[T]{a: s.Specification}
	ns.Associate(ns)
	return ns
}

func (s *Spec[T]) Or(other Specification[T]) Specification[T] {
	ns := &Or[T]{a: s.Specification, b: other}
	ns.Associate(ns)
	return ns
}

func (s *Spec[T]) Associate(spec Specification[T]) {
	s.Specification = spec
}
