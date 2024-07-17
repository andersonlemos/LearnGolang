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

func checkIfContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, obj := range haystack {
		if obj == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("Expected %s to contain %s", needle, haystack)
	}
}
func TestWalks(t *testing.T) {
	t.Run("walks on structs", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			CallsExpected []string
		}{
			{
				"Struct with string fields",
				struct {
					Name string
				}{"Chris"},
				[]string{"Chris"},
			},
			{
				"struct with two fields",
				struct {
					Name string
					City string
				}{"Chris", "London"},
				[]string{"Chris", "London"},
			},
			{
				"Struct sem campo tipo string",
				struct {
					Nome  string
					Idade int
				}{"Chris", 33},
				[]string{"Chris"},
			},
			{
				"Concatenated fields",
				Person{
					"Chris",
					Profile{
						33,
						"London",
					},
				},
				[]string{"Chris", "London"},
			},
			{
				"Pointer of things",
				&Person{
					"Chris",
					Profile{33, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"Slices",
				[]Profile{
					{33, "London"},
					{34, "Reykjavík"},
				},
				[]string{"London", "Reykjavík"},
			},
		}

		for _, c := range cases {
			t.Run(c.Name, func(t *testing.T) {
				var result []string
				Walk(c.Input, func(input string) {
					result = append(result, input)
				})
				if !reflect.DeepEqual(result, c.CallsExpected) {
					t.Errorf("got %v, want %v", result, c.CallsExpected)
				}
			})
		}
	})
	t.Run("walks on maps", func(t *testing.T) {
		mapA := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}
		var result []string
		Walk(mapA, func(input string) {
			result = append(result, input)
		})
		checkIfContains(t, result, "Bar")
		checkIfContains(t, result, "Boz")
	})
}
