package convert

import (
	"fmt"
	"testing"
)

func TestToJson(t *testing.T) {
	t.Parallel()

}

func TestToMap(t *testing.T) {
	t.Parallel()

	type Message struct {
		ID      int
		Name    string
		Content string
	}
	messages := []Message{
		{ID: 1, Name: "test1", Content: "test1 content"},
		{ID: 2, Name: "test2", Content: "test2 content"},
	}

	result := ToMap(messages, func(msg Message) (int, Message) {
		return msg.ID, msg
	})

	fmt.Println(result)
}
