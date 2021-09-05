// Package utils
package utils

import "testing"

func TestReplacer(t *testing.T) {
	t.Parallel()
	testCase := map[string]string{
		"testCase#1": "testCase_1",
		"testCase@1": "testCase-1",
	}

	rule := map[string]string{
		"#": "_",
		"@": "-",
	}

	for k, v := range testCase {
		ns := StringReplacer(k, rule)
		if v == ns {

			t.Logf(`expected: %v, got: %v`, v, StringReplacer(k, rule))
		}
	}
}
