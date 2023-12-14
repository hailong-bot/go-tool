package stream

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	t.Parallel()
	type Person struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}

	needs := []Person{
		{Name: "Tom", Age: 18},
		{Name: "Mike", Age: 19},
	}
	res := Of(needs...).Filter(func(p Person) bool {
		return p.Age > 18
	}).Map(func(p Person) Person {
		p.Age++
		p.Name = "change Name"
		return p
	}).Collect(&ToCollection[Person]{})
	fmt.Println("res", res)
}
