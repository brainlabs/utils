// Package utils
// @author Daud Valentino
package utils

import "testing"

var testCase = []struct {
	Value   interface{}
	IsEmpty bool
}{
	{
		"", true,
	},
	{
		" ", true,
	},
	{
		" jakarta", false,
	},
	{
		int64(0), true,
	},
	{
		uint(0), true,
	},
	{
		float64(0), true,
	},
	{
		[]string{}, true,
	},
	{
		map[string]interface{}{}, true,
	},
	{
		false, true,
	},
	{
		nil, true,
	},
	{
		struct {
		}{}, false,
	},
	{
		new(struct{}), false,
	},
}

func TestIsEmptyValue(t *testing.T) {

	for _, c := range testCase {
		ie := IsEmptyValue(c.Value)
		if ie == c.IsEmpty {
			t.Logf("expected %v, got actual: %v", c.IsEmpty, ie)
			continue
		}

		t.Errorf("expected %v, got actual: %v", c.IsEmpty, ie)
	}
}
