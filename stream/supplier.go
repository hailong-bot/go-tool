package stream

// Supplier Represents a supplier of results.
type Supplier[T any] interface {
	get() T
}
