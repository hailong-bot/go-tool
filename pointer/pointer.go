// Pacakge pointer contains some util function to operator go pointer
package pointer

import "reflect"

// Of returns a pointer to the value 'v'
func Of[T any](v T) *T {
	return &v
}

// Unwrap returns thr value from the pointer
func Unwarp[T any](p *T) T {
	return *p
}

// UnwrapOr returns the value from to the pointer or fallback if the pointer is nil
func UnwarpOr[T any](p *T, fallback T) T {
	if p == nil {
		return fallback
	}

	return *p
}

// UnwrapOrDefault returns the value from the pointer or the default value if the pointer is nil
func UnwarpOrDefault[T any](p *T) T {
	var v T
	if p == nil {
		return v
	}
	return *p
}

// ExtractPointer returns ther underlying by the given interface type
func ExtractPointer(value any) any {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	if t.Kind() != reflect.Pointer {
		return value
	}
	return ExtractPointer(v.Elem().Interface())
}
