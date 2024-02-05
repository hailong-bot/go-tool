package structs

import (
	"errors"
	"reflect"

	"github.com/hailong-bot/go-tool/v2/pointer"
)

// defaultTagName is the default tag for struct fields to lookup
var defaultTagName = "json"

type Struct struct {
	raw     any
	rtype   reflect.Type
	rvalue  reflect.Value
	TagName string
}

// New returns a new *Struct
func New(value any, tagName ...string) *Struct {
	value = pointer.ExtractPointer(value)
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	tn := defaultTagName
	if len(tagName) > 0 {
		tn = tagName[0]
	}

	return &Struct{
		raw:     value,
		rtype:   t,
		rvalue:  v,
		TagName: tn,
	}
}

// ToMap converts the given struct to a map[string]any, where the keys
// of the keys are the field names and the values of the map are the values
// of the fields. The default map key is the struct field name, but you can
// change it. The `json` key is the default tag key. Example:
//
//	 // default
//	 Name string `json:"name"`
//
//
//	 // ignore the fields
//	 Name string // no tag
//		Age  string `json:"-"` // json ignore tag
//		sex  string // unexported field
//		Goal int    `json:"goal,omitempty"` // omitempty if the field is zero value
//
// ToMap convert the exported fields of a struct to map.
//
// //custom map key
// Name string `json:"myName"`
func (s *Struct) ToMap() (map[string]any, error) {
	if !s.IsStruct() {
		return nil, errors.New("value is not struct")
	}

	result := make(map[string]any)
	fields := s.Fields()
	for _, f := range fields {
		if !f.IsExported() || f.tag.IsEmpty() || f.tag.Name == "-" {
			continue
		}
		if f.IsZero() && f.tag.HasOption("omitempty") {
			continue
		}
		result[f.tag.Name] = f.mapValue(f.Value())
	}

	return result, nil
}

// Fields returns all the struct fields within a slice
func (s *Struct) Fields() []*Field {
	var fields []*Field
	fieldNum := s.rvalue.NumField()
	for i := 0; i < fieldNum; i++ {
		v := s.rvalue.Field(i)
		sf := s.rtype.Field(i)
		field := newField(v, sf, s.TagName)
		fields = append(fields, field)
	}
	return fields
}

// Field returns a Field if the given field name was found
func (s *Struct) Field(name string) (*Field, bool) {
	f, ok := s.rtype.FieldByName(name)
	if !ok {
		return nil, false
	}
	return newField(s.rvalue.FieldByName(name), f, s.TagName), true
}

func (s *Struct) IsStruct() bool {
	k := s.rvalue.Kind()
	if k == reflect.Invalid {
		return false
	}
	return k == reflect.Struct
}

// ToMap convert struct tp map, only convert exported struct field
// map key is specified same as struct field tag `json` value
func ToMap(v any) (map[string]any, error) {
	return New(v).ToMap()
}
