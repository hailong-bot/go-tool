package structs

import (
	"strings"

	"github.com/hailong-bot/go-tool/v2/validator"
)

// Tag is abstact struct field tag
type Tag struct {
	Name    string
	Options []string
}

func newTag(tag string) *Tag {
	res := strings.Split(tag, ",")
	return &Tag{
		Name:    res[0],
		Options: res[1:],
	}
}

func (t *Tag) HasOption(opt string) bool {
	for _, o := range t.Options {
		if o == opt {
			return true
		}
	}
	return false
}

func (t *Tag) IsEmpty() bool {
	return validator.IsEmptyString(t.Name)
}
