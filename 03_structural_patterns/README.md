# Structural Patterns

## Composite Pattern

### Why is there no test?

It doesn't have much to test apart from the structure itself.
We won't write unit tests, and we'll simply describe the ways to create those compositions in Go.

### Composite pattern versus inheritance

In go, you can simply composite the son struct (sub class in the object oriented programming) with the parent without embedding so that you can access the Parent instance later:

```go
type Parent struct {
    SomeField int
}

type Son struct {
    P Parent
}

func GetParentField(p *Parent) int {
    fmt.Println(p.SomeField)
}

son := Son{}
GetParentField(son.P)
```
