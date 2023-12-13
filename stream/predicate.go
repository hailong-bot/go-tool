package stream

// Predicate Represents a predicate (boolean-valued function) of one argument.
type Predicate[T any] interface {
	test(t T) bool
}
