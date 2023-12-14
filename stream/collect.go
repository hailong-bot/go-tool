package stream

type Collector[T any] interface {
	Accumulate(T)
	Result() T
}

type ToCollection[T any] struct {
	slice []T
}

func (c *ToCollection[T]) Accumulate(item T) {
	c.slice = append(c.slice, item)
}

func (c *ToCollection[T]) Result() []T {
	return c.slice
}
