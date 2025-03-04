# Structural Patterns

## Composite Pattern

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
