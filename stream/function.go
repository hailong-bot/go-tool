package stream

// Function Represents a function that accepts one argument and produces a result.
type Function[T any, R any] interface {
	apply(t T) R
}
