package guerystring

import (
	"testing"
)

func assert(t *testing.T, x interface{}, expectedQuery string, expectedErr error) {
	actual, err := Encode(x)
	if expectedErr != nil {
		if err.Error() != expectedErr.Error() {
			t.Logf("expect err is '%v' but got '%v'", expectedErr, err)
			t.Fail()
		}
	} else {
		if expectedQuery != actual {
			t.Logf("expect query is '%v' but got '%v'", expectedQuery, actual)
			t.Fail()
		}
	}
}

func TestEncode(t *testing.T) {
	assert(t,
		true,
		"",
		errNotStruct,
	)
	assert(t,
		1,
		"",
		errNotStruct,
	)
	assert(t,
		"not a struct",
		"",
		errNotStruct,
	)
	assert(t,
		map[string]string{},
		"",
		errNotStruct,
	)
	assert(t,
		[]string{"string1", "string2"},
		"",
		errNotStruct,
	)
	assert(t,
		struct {
			String1 string `query:"string1"`
			String2 string `query:"string2"`
			String3 string `query:"string3"`
		}{
			String1: "string1",
			String2: "string2",
			String3: "string3",
		},
		"string1=string1&string2=string2&string3=string3",
		nil,
	)
	assert(t,
		struct {
			String1 string `query:"string1"`
			String2 string
			String3 string `query:"string3"`
		}{
			String1: "string1",
			String2: "string2",
			String3: "string3",
		},
		"string1=string1&string3=string3",
		nil,
	)
	assert(t,
		struct {
			Boolean bool   `query:"boolean"`
			Integer int    `query:"integer"`
			String1 string `query:"string1"`
			String2 string `query:"string2"`
		}{
			Boolean: true,
			Integer: 1,
			String1: "string1",
			String2: "string2",
		},
		"string1=string1&string2=string2",
		nil,
	)
}
