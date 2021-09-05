// Package utils
package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestToString(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		Input    interface{}
		Expected string
	}{
		{
			Input:    int(1),
			Expected: "1",
		},
		{
			Input:    int8(2),
			Expected: "2",
		},
		{
			Input:    int16(3),
			Expected: "3",
		},
		{
			Input:    int32(4),
			Expected: "4",
		},
		{
			Input:    int64(5),
			Expected: "5",
		},
		{
			Input:    uint(6),
			Expected: "6",
		},
		{
			Input:    uint8(7),
			Expected: "7",
		},
		{
			Input:    uint16(8),
			Expected: "8",
		},
		{
			Input:    uint32(9),
			Expected: "9",
		},
		{
			Input:    uint64(10),
			Expected: "10",
		},
		{
			Input:    float32(11),
			Expected: "11",
		},
		{
			Input:    float64(12),
			Expected: "12",
		},
		{
			Input:    bool(true),
			Expected: "true",
		},
		{
			Input: struct {
			}{},
			Expected: "{}",
		},

		{
			Input:    string(`test`),
			Expected: "test",
		},
	}

	for _, sc := range scenarios {
		r := ToString(sc.Input)

		if sc.Expected == r {
			t.Logf(`expected: %v, got: %v`, sc.Expected, r)
			continue
		}

		t.Errorf(`expected: %v, got: %v`, sc.Expected, r)
	}
}

func TestStructToMap(t *testing.T) {
	type (
		Contoh struct {
			ID        int64     `json:"id" store:"id"`
			Name      string    `json:"name" store:"name,omitempty"`
			CreatedAt time.Time `json:"created_at"`
			alamat    string
		}
	)

	testCase := []struct {
		Input    interface{}
		CountCol int
		Error    error
	}{
		{
			Input: Contoh{
				ID:        1,
				Name:      "test",
				CreatedAt: time.Now(),
				alamat:    "test",
			},
			CountCol: 2,
			Error:    nil,
		},
		{
			Input: &Contoh{
				ID:        1,
				Name:      "",
				CreatedAt: time.Now(),
				alamat:    "test",
			},
			CountCol: 1,
			Error:    nil,
		},
		{
			Input: map[string]interface{}{
				"name": "test 2",
			},
			CountCol: 0,
			Error:    fmt.Errorf("only accepted struct, got map"),
		},
		{
			Input:    "test",
			CountCol: 0,
			Error:    fmt.Errorf("only accepted struct, got map"),
		},
	}

	for _, x := range testCase {
		c, e := StructToMap(x.Input, "store")
		if fmt.Sprintf("%T", x.Error) == fmt.Sprintf("%T", e) {

			t.Logf("expected '%v', but got '%v'", x.Error, e)
		} else {
			t.Errorf("expected '%v', but got '%v'", x.Error, e)
		}

		if len(c) == x.CountCol {
			t.Logf("expected %v, but got %v", x.CountCol, len(c))
		} else {
			t.Errorf("expected %v, but got %v", x.CountCol, len(c))
		}


	}
}

func TestToColumnsValues(t *testing.T) {

	type (
		Contoh struct {
			ID        int64     `json:"id" store:"id"`
			Name      string    `json:"name" store:"name,omitempty"`
			CreatedAt time.Time `json:"created_at"`
			alamat    string
		}
	)

	testCase := []struct {
		Input    interface{}
		CountCol int
		Error    error
	}{
		{
			Input: Contoh{
				ID:        1,
				Name:      "test",
				CreatedAt: time.Now(),
				alamat:    "test",
			},
			CountCol: 2,
			Error:    nil,
		},
		{
			Input: &Contoh{
				ID:        1,
				Name:      "",
				CreatedAt: time.Now(),
				alamat:    "test",
			},
			CountCol: 1,
			Error:    nil,
		},
		{
			Input: map[string]interface{}{
				"name": "test 2",
			},
			CountCol: 0,
			Error:    fmt.Errorf("only accepted struct, got map"),
		},
		{
			Input:    "test",
			CountCol: 0,
			Error:    fmt.Errorf("only accepted struct, got map"),
		},
	}

	for _, x := range testCase {
		c, v, e := StructExtractFieldValue(x.Input, "store")
		if fmt.Sprintf("%T", x.Error) == fmt.Sprintf("%T", e) {

			t.Logf("expected '%v', but got '%v'", x.Error, e)
		} else {
			t.Errorf("expected '%v', but got '%v'", x.Error, e)
		}

		if len(c) == x.CountCol {
			t.Logf("expected %v, but got %v", x.CountCol, len(c))
		} else {
			t.Errorf("expected %v, but got %v", x.CountCol, len(c))
		}

		if len(v) == x.CountCol {
			t.Logf("expected %v, but got %v", x.CountCol, len(v))
		} else {
			t.Errorf("expected %v, but got %v", x.CountCol, len(v))
		}

	}

}
