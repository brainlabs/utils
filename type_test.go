// Package utils
package utils

import "testing"

func Test_IsSamType(t *testing.T) {

	cases := []struct {
		Source      interface{}
		Destination interface{}
		IsSame      bool
	}{
		{
			Source:      nil,
			Destination: nil,
			IsSame:      true,
		},
		{
			Source:      struct{}{},
			Destination: new(string),
			IsSame:      false,
		},
	}

	for _, c := range cases {
		ie := IsSameType(c.Source, c.Destination)
		if ie == c.IsSame {
			t.Logf("expected %T, got actual: %T, should be : %v", c.Source, c.Destination, c.IsSame)
			continue
		}

		t.Errorf("expected %T, got actual: %T, should be : %v", c.Source, c.Destination, c.IsSame)
	}
}
