package specification

import "context"

// Specification interface.
// Use BaseSpecification as base for creating specifications, and
// only the method predicate(Object) must be implemented.
type Specification[T any] interface {
	// IsSatisfiedBy check if t is satisfied by the specification.
	IsSatisfiedBy(ctx context.Context, t T) bool
	// And create a new specification that is the AND operation of the current specification and
	// another specification.
	And(another Specification[T]) Specification[T]
	// Or create a new specification that is the OR operation of the current specification and
	// another specification.
	Or(another Specification[T]) Specification[T]
	// Not create a new specification that is the NOT operation of the current specification.
	Not(another Specification[T]) Specification[T]
}

type BaseSpecification[T any] struct {
	predicate func(ctx context.Context, t T) bool
}

func New[T any](predicate func(ctx context.Context, t T) bool) Specification[T] {
	return &BaseSpecification[T]{predicate: predicate}
}

func (spec *BaseSpecification[T]) IsSatisfiedBy(ctx context.Context, t T) bool {
	return spec.predicate(ctx, t)
}

func (spec *BaseSpecification[T]) And(another Specification[T]) Specification[T] {
	return And[T](spec, another)
}

func (spec *BaseSpecification[T]) Or(another Specification[T]) Specification[T] {
	return Or[T](spec, another)
}

func (spec *BaseSpecification[T]) Not(another Specification[T]) Specification[T] {
	return Not[T](another)
}
