package stream

import (
	"fmt"
	"testing"

	"github.com/hailong-bot/go-tool/v2/convert"
)

func Test1(t *testing.T) {
	t.Parallel()
	type Person struct {
		ID   int64  `json:"ID"`
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}

	needs := []Person{
		{ID: 1, Name: "Tom", Age: 18},
		{ID: 2, Name: "Mike", Age: 19},
		{ID: 3, Name: "Mose", Age: 20},
		{ID: 1, Name: "Tom", Age: 18},
	}
	res := Of(needs...).Filter(func(p Person) bool {
		return p.Age > 17
	}).Map(func(p Person) Person {
		if p.ID == 3 {
			p.Age++
			p.Name = "change Name"
		}
		return p
	}).Distinct().Collect(&ToCollection{})

	var result []Person
	err := convert.Convert(res, &result)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("res", result)
}
