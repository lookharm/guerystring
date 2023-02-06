# guerystring

The `guerystring` package can encode Go struct into URL query strings.
Now only the string type is supported.

## Example
```go
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
```