// Package utils
package utils

import "reflect"

// IsSameType check type from source to destination
func IsSameType(src, dest interface{}) bool  {
	return reflect.TypeOf(src) == reflect.TypeOf(dest)
}
