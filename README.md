# guerystring

The `guerystring` package can encode Go struct into URL query strings.
Now only the string type is supported.

## Usage
```go
import "github.com/lookharm/guerystring"
```

## Example
```go
type Request struct {
    String1 string `query:"string1"`
    String2 string `query:"string2"`
    String3 string `query:"string3"`
}

r := Request{
    String1: "string1",
    String2: "string2",
    String3: "string3",
}

query, err := Encode(r) // string1=string1&string2=string2&string3=string3
```