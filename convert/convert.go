package convert

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// Play:
// ToJsonString convert value to json string
func ToJsonString(value any) (string, error) {
	result, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// Play:
// Convert convert source value to destination value based on type
func Convert(source any, dest any) error {
	sourceValue := reflect.ValueOf(source)
	targetValue := reflect.ValueOf(dest)

	if targetValue.Kind() != reflect.Ptr {
		return fmt.Errorf("target must be a pointer")
	}

	if sourceValue.Kind() == reflect.Slice {
		if targetValue.Elem().Kind() != reflect.Slice {
			return fmt.Errorf("source is a slice, but target is not a slice")
		}

		targetSliceType := reflect.SliceOf(targetValue.Elem().Type().Elem())
		targetSlice := reflect.MakeSlice(targetSliceType, 0, sourceValue.Len())

		for i := 0; i < sourceValue.Len(); i++ {
			elem := reflect.New(targetSliceType.Elem()).Interface()
			err := Convert(sourceValue.Index(i).Interface(), elem)
			if err != nil {
				return err
			}
			targetSlice = reflect.Append(targetSlice, reflect.ValueOf(elem).Elem())
		}

		reflect.ValueOf(dest).Elem().Set(targetSlice)
		return nil
	}

	if sourceValue.Kind() == reflect.Map {
		if targetValue.Elem().Kind() != reflect.Map {
			return errors.New("source is a map, but target is not a map")
		}
		targetMapType := reflect.MapOf(targetValue.Elem().Type().Key(), targetValue.Elem().Type().Elem())
		targetMap := reflect.MakeMap(targetMapType)

		for _, key := range sourceValue.MapKeys() {
			elem := reflect.New(targetMapType.Elem()).Interface()
			err := Convert(sourceValue.MapIndex(key).Interface(), elem)
			if err != nil {
				return err
			}
			targetMap.SetMapIndex(key, reflect.ValueOf(elem).Elem())
		}

		reflect.ValueOf(dest).Elem().Set(targetMap)

		return nil
	}

	if !sourceValue.Type().ConvertibleTo(targetValue.Elem().Type()) {
		return fmt.Errorf("cannot convert source to target type")
	}

	convertedValue := sourceValue.Convert(targetValue.Elem().Type())
	reflect.ValueOf(dest).Elem().Set(convertedValue)

	return nil
}

// Play:
// ToInt64 convert value to int64 value , if input is not numerical, return 0 and err
func ToInt64(value any) (int64, error) {
	v := reflect.ValueOf(value)

	var result int64
	err := fmt.Errorf("ToInt64: invalid value type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		result = v.Int()
		return result, nil
	case uint, uint8, uint16, uint32, uint64:
		result = int64(v.Uint())
		return result, nil
	case float32, float64:
		result = int64(v.Float())
		return result, nil
	case string:
		result, err = strconv.ParseInt(v.String(), 0, 64)
		if err != nil {
			result = 0
		}
		return result, err
	default:
		return result, err
	}
}

// Play:
// ToMap convert a slice of structs to a map based on interatee function.
func ToMap[T any, K comparable, V any](array []T, iteratee func(T) (K, V)) map[K]V {
	result := make(map[K]V, len(array))
	for _, item := range array {
		k, v := iteratee(item)
		result[k] = v
	}
	return result
}
