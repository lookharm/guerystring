package guerystring

import (
	"testing"
)

func assertEqual(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Logf("expect %v but got %v", expected, actual)
		t.Fail()
	}
}

func TestEncode(t *testing.T) {
	type Request struct {
		String1 string `query:"string1"`
		String2 string `query:"string2"`
		String3 string `query:"string3"`
	}

	r := Request{
		String1: "1",
		String2: "2",
		String3: "3",
	}
	query, err := Encode(r)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, `string1=1&string2=2&string3=3`, query)
}
