package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

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
			{
				testName: "struct with two string fields",
				input: struct {
					Name string
					City string
				}{Name: "Chris", City: "London"},
				calls: []string{"Chris", "London"},
			},
			{
				testName: "struct with string and int",
				input: struct {
					Name string
					Age  int
				}{Name: "Chris", Age: 21},
				calls: []string{"Chris"},
			},
			{
				testName: "nested struct",
				input: Person{
					Name:    "Chris",
					Profile: Profile{Age: 18, City: "London"},
				},
				calls: []string{"Chris", "London"},
			},
			{
				testName: "pointer to things",
				input: &Person{
					Name:    "Chris",
					Profile: Profile{Age: 18, City: "London"},
				},
				calls: []string{"Chris", "London"},
			},
			{
				testName: "slices",
				input: []Profile{
					{18, "London"},
					{23, "Berlin"},
				},
				calls: []string{"London", "Berlin"},
			},
			{
				testName: "arrays",
				input: [2]Profile{
					{18, "London"},
					{23, "Berlin"},
				},
				calls: []string{"London", "Berlin"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"foo": "bar",
			"boz": "baz",
		}

		var got []string
		walk(aMap, func(s string) {
			got = append(got, s)
		})

		assertContains(t, got, "bar")
		assertContains(t, got, "baz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
