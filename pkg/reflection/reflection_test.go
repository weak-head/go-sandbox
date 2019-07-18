package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	t.Run("fixed test cases", func(t *testing.T) {
		testCases := []struct {
			testName string
			input    interface{}
			calls    []string
		}{
			{
				testName: "single string field",
				input:    struct{ Name string }{"Chris"},
				calls:    []string{"Chris"},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.testName, func(t *testing.T) {
				var got []string
				walk(tc.input, func(s string) {
					got = append(got, s)
				})
				if !reflect.DeepEqual(got, tc.calls) {
					t.Errorf("got %s, want %s", got, tc.calls)
				}
			})
		}
	})
}
