package walker

import (
	"reflect"
	"testing"
)

func TestWalker(t *testing.T) {

	t.Run("Table tests", func(t *testing.T) {
		type TestCase struct {
			Name   string
			Input  any
			Output []string
		}

		var tt []TestCase = []TestCase{
			{"empty struct", struct{}{}, []string{}},
			{"struct with int field", struct{ int }{1}, []string{}},
			{"struct with 2 fields", struct {
				A string
				B string
			}{"a", "b"}, []string{"a", "b"}},
			{"nested struct", struct {
				A string
				B struct{ C string }
			}{"a", struct{ C string }{"c"}}, []string{"a", "c"}},
			{"pointer to things", &struct {
				A string
				B struct{ C string }
			}{"a", struct{ C string }{"c"}}, []string{"a", "c"}},

			{
				"slices",
				[]struct{ A string }{
					{"a"},
					{"b"},
				},
				[]string{"a", "b"},
			},

			{"arrays",
				[2]struct {
					int
					string
				}{
					{33, "London"},
					{34, "Reykjavík"},
				},
				[]string{"London", "Reykjavík"},
			},

			{
				"maps",
				map[string]string{
					"Cow":   "Moo",
					"Sheep": "Baa",
				},
				[]string{"Moo", "Baa"},
			},
		}

		for _, i := range tt {
			t.Run(i.Name, func(t *testing.T) {
				want := i.Output
				got := make([]string, 0, cap(i.Output))

				Walker(i.Input, func(s string) {
					got = append(got, s)
				})

				if !reflect.DeepEqual(want, got) {
					t.Errorf("want -> %q, %q <- got", want, got)
				}
			})
		}
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan struct {
			int
			string
		})

		go func() {
			aChannel <- struct {
				int
				string
			}{33, "Berlin"}
			aChannel <- struct {
				int
				string
			}{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walker(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
