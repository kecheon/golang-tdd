package reflection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	t.Run("Common types", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"Struct with 1 string field",
				struct {
					Name string
				}{"Chris"},
				[]string{"Chris"},
			},
			{
				"Struct with 2 string fields",
				struct {
					Name string
					City string
				}{"Chris", "London"},
				[]string{"Chris", "London"},
			},
			{
				"Struct with 1 string field and 1 int field",
				struct {
					Name string
					Age  int
				}{"Chris", 61},
				[]string{"Chris"},
			},
			{
				"Nested fields",
				struct {
					Name    string
					Profile struct {
						Age  int
						City string
					}
				}{
					"Chris", struct {
						Age  int
						City string
					}{33, "London"}},
				[]string{"Chris", "London"},
			},
			{
				"Pointers to things",
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
					{61, "Oya"},
				},
				[]string{"London", "Oya"},
			},
		}
		for _, c := range cases {
			t.Run(c.Name, func(t *testing.T) {
				var got []string
				Walk(c.Input, func(input string) {
					got = append(got, input)
				})
				if !reflect.DeepEqual(got, c.ExpectedCalls) {
					t.Errorf("got %v, Expected %v", got, c.ExpectedCalls)
				}
			})
		}
	})

	t.Run("Maps don't care order", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("Channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{33, "London"}
			aChannel <- Profile{61, "서울"}
			close(aChannel)
		}()
		var got []string
		want := []string{"London", "서울"}
		Walk(aChannel, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Function", func(t *testing.T) {
		aFunc := func() (Profile, Profile) {
			return Profile{33, "London"}, Profile{61, "서울"}
		}
		var got []string
		want := []string{"London", "서울"}
		Walk(aFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			fmt.Printf("x: %s, needle: %s, contains: %t\n", x, needle, contains)
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
