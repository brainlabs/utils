// Package utils
package utils

import "testing"

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestGenerateRandomNumberString(t *testing.T) {
	t.Parallel()
	t.Log("Test generation of string number")
	{
		var tester []string
		for i := 0; i < 20000; i++ {
			ref := GenerateRandomNumberString(6)
			tester = append(tester, ref)
		}

		var isSame bool

		for i := 0; i < len(tester); i++ {

			for j := 0; j < len(tester); j++ {

				if i != j {
					isSame = tester[i] == tester[j]
				}

			}

		}

		if isSame {
			t.Errorf("%s expected no two numbers have the same value", failed)
		} else {
			t.Logf("%s expected no two numbers have the same value", success)
		}
	}
}
