package stream

type Collector interface {
	Accumulate(any)
	Result() any
}

type ToCollection struct {
	slice []interface{}
}

func (c *ToCollection) Accumulate(item any) {
	c.slice = append(c.slice, item)
}

func (c *ToCollection) Result() any {
	return c.slice
}
