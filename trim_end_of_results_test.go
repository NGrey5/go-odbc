package odbc

import "testing"

type TestStruct struct {
	Hello   string
	Int     int
	String2 string
	Float   float64
}

func TestTrimEndOfResults(t *testing.T) {
	s := TestStruct{
		Hello:   "hello    ",
		Int:     10,
		String2: "     string2",
		Float:   0.4,
	}

	TrimEndOfResults(&s)

	want := TestStruct{
		Hello:   "hello",
		Int:     10,
		String2: "string2",
		Float:   0.4,
	}

	if want != s {
		t.Fatalf("Wanted %+v, Got %+v", want, s)
	}
}

func TestTrimEndOfResultsSlice(t *testing.T) {
	s := []TestStruct{
		{
			Hello:   "hello1    ",
			Int:     10,
			String2: "     string1",
			Float:   0.4,
		},
		{
			Hello:   "hello2    ",
			Int:     20,
			String2: "     string2",
			Float:   22.5,
		},
	}

	TrimEndOfResults(&s)

	want := []TestStruct{
		{
			Hello:   "hello1",
			Int:     10,
			String2: "string1",
			Float:   0.4,
		},
		{
			Hello:   "hello2",
			Int:     20,
			String2: "string2",
			Float:   22.5,
		},
	}

	for i := range want {
		if want[i] != s[i] {
			t.Fatalf("Wanted %v, Got %+v", want[i], s[i])
		}
	}

}
