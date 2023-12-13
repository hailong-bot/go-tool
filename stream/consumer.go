package stream

// Consumer Represents an operation that accepts a single input argument and returns no
// result. Unlike most other functional interfaces, {@code Consumer} is expected
// to operate via side-effects.
type Consumer[T any] interface {
	accept(t T)
}
