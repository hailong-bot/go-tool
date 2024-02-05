package convert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/hailong-bot/go-tool/v2/structs"
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

// ToString convert value to string
func ToString(value any) string {
	if value == nil {
		return ""
	}

	switch val := value.(type) {
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case int:
		return strconv.FormatInt(int64(val), 10)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case string:
		return val
	case []byte:
		return string(val)
	default:
		b, err := json.Marshal(val)
		if err != nil {
			return ""
		}
		return string(b)

		// todo: maybe we should't supprt other type conversion
		// v := reflect.ValueOf(value)
		// log.Panicf("Unsupported data type: %s ", v.String())
		// return ""
	}
}

// ToFloat convet value to floate64
func ToFloat64(value any) (float64, error) {
	v := reflect.ValueOf(value)

	result := 0.0
	err := fmt.Errorf("ToInt: unvalid interface type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		result = float64(v.Int())
		return result, nil
	case uint, uint8, uint16, uint32, uint64:
		result = float64(v.Uint())
		return result, nil
	case float32, float64:
		result = v.Float()
		return result, nil
	case string:
		result, err = strconv.ParseFloat(v.String(), 64)
		if err != nil {
			result = 0.0
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

// Play:
// StructToMap  convert struct to map, only convert exported struct field
func StructToMap(value any) (map[string]any, error) {
	return structs.ToMap(value)
}
