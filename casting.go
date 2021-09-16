// Package utils
package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ToString converts a value to string.
func ToString(value interface{}) string {
	switch value := value.(type) {
	case string:
		return value
	case int:
		return strconv.FormatInt(int64(value), 10)
	case int8:
		return strconv.FormatInt(int64(value), 10)
	case int16:
		return strconv.FormatInt(int64(value), 10)
	case int32:
		return strconv.FormatInt(int64(value), 10)
	case int64:
		return strconv.FormatInt(int64(value), 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(uint64(value), 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case float64:
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	}

	return fmt.Sprintf("%+v", value)
}


// ToMap converts a struct to a map using the struct's tags.
//
// ToMap uses tags on struct fields to decide which fields to add to the
// returned map.
func StructToMap(src interface{}, tag string) (map[string]interface{}, error) {
	out := map[string]interface{}{}
	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return out, fmt.Errorf("only accepted %s, got %s", reflect.Struct.String(), v.Kind().String())
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)

		tagsv := strings.Split(fi.Tag.Get(tag), ",")

		if tagsv[0] != "" && fi.PkgPath == "" {
			// skip if omitempty
			if (len(tagsv) > 1 && tagsv[1] == "omitempty") && IsEmptyValue(v.Field(i).Interface()) {
				continue
			}

			if isTime(v.Field(i)) {
				if v.Field(i).Interface().(time.Time).IsZero() && tagsv[1] == "omitempty" {
					continue
				}
			}

			// set key value of map interface output
			out[tagsv[0]] = v.Field(i).Interface()

		}
	}

	return out, nil
}

// StructExtractFieldValue iterate struct to separate key field and value
func StructExtractFieldValue(src interface{}, tag string) ([]string, []interface{}, error) {
	var columns []string
	var values []interface{}

	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, nil, fmt.Errorf("only accepted %s, got %s", reflect.Struct.String(), v.Kind().String())
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)


		tagsv := strings.Split(fi.Tag.Get(tag), ",")

		if tagsv[0] != "" && fi.PkgPath == "" {
			// skip if omitempty
			if (len(tagsv) > 1 && tagsv[1] == "omitempty") && IsEmptyValue(v.Field(i).Interface()) {
				continue
			}

			if isTime(v.Field(i)) {
				if v.Field(i).Interface().(time.Time).IsZero() && tagsv[1] == "omitempty" {
					continue
				}
			}

			// set value of string slice to value in struct field
			columns = append(columns, tagsv[0])

			// set value interface of value struct field
			values = append(values, v.Field(i).Interface())

		}
	}

	return columns, values, nil
}

func StructToBulkInsert(src interface{}, tag string) ([]string, []interface{}, []string, error) {
	var columns []string
	var replacers []string
	var values []interface{}

	v := reflect.Indirect(reflect.ValueOf(src))
	t := reflect.TypeOf(src)
	if t.Kind() == reflect.Ptr {
		//v = v.Elem()
		t = t.Elem()
	}

	if t.Kind() == reflect.Slice {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return columns, values, replacers, fmt.Errorf("only accepted %s, got %s", reflect.Struct.String(), t.Kind().String())
	}

	for i := 0; i < v.Len(); i++ {

		item := v.Index(i)
		if !item.IsValid() {
			continue
		}

		cols, val, err := StructExtractFieldValue(item.Interface(), tag)
		if err != nil {
			return columns, values, replacers, err
		}

		if len(columns) == 0 {
			columns = cols
		}

		pattern := fmt.Sprintf(`(%s)`, strings.TrimRight(strings.Repeat("?,", len(columns)), `,`))
		replacers = append(replacers, pattern)
		values = append(values, val...)
	}

	return columns, values, replacers, nil

}

func isTime(obj reflect.Value) bool {
	_, ok := obj.Interface().(time.Time)
	return ok
}
